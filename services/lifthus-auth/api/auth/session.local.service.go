package auth

import (
	"encoding/json"
	"fmt"
	"lifthus-auth/common/db"
	"lifthus-auth/common/helper"
	"lifthus-auth/common/lifthus"
	"strconv"
	"strings"

	"lifthus-auth/common/service/session"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// newSessionHandler is a local development version of NewSessionHandler.
// which uses Authorization header instead of cookie.
func (ac authApiController) newSessionHandler(c echo.Context) error {
	// get token from Authorization header
	authorizationHeader := c.Request().Header.Get("Authorization")

	var sid, stSigned, uid string
	var exp bool
	var err error

	// case A: no session, create new session
	if authorizationHeader == "" {
		// create new session
		sid, stSigned, err = session.CreateSession(c.Request().Context(), ac.dbClient)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, err.Error())
		}
		// case B,C: if session token exists, validate it.
	} else {
		if !strings.HasPrefix(authorizationHeader, "Bearer ") {
			return c.String(http.StatusBadRequest, "invalid authorization header")
		}
		lifthus_st := authorizationHeader[7:]

		// ValidateSessionToken validates the token for case B,C, and updates the db for case B-1.
		sid, uid, exp, err = session.ValidateSession(c.Request().Context(), ac.dbClient, lifthus_st)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, err.Error())

			// case B-1,2: if it is expired, refresh the token using same SID but clear the UID.
		} else if exp {
			stSigned, err = session.RefreshSessionToken(c.Request().Context(), ac.dbClient, sid)
			if err != nil {
				log.Println(err)
				return c.String(http.StatusInternalServerError, err.Error())
			}

			// case C: if it is valid, just keep the token and return
		} else {
			if uid != "" { // case C-1
				// get user by uid
				uidUint64, err := strconv.ParseUint(uid, 10, 64)
				if err != nil {
					log.Println(err)
					return c.String(http.StatusInternalServerError, err.Error())
				}
				ls, err := db.QueryUserByUID(c.Request().Context(), ac.dbClient, uidUint64)
				if err != nil {
					log.Println(err)
					return c.String(http.StatusInternalServerError, err.Error())
				}
				// make struct with UID and Name
				keepResp := struct {
					UID  string `json:"uid"`
					Name string `json:"username"`
				}{
					UID:  strconv.FormatUint(ls.ID, 10),
					Name: ls.Name,
				}
				keepRespJSON, err := json.Marshal(keepResp)
				if err != nil {
					log.Println(err)
					return c.String(http.StatusInternalServerError, err.Error())
				}
				c.Response().Header().Set("Authorization", "Bearer "+lifthus_st)
				return c.JSONBlob(http.StatusOK, keepRespJSON)
			} else { // case C-2

				c.Response().Header().Set("Authorization", "Bearer "+lifthus_st)
				return c.String(http.StatusCreated, sid)
			}
		}
	}

	// set session token in cookie.
	cookie := &http.Cookie{
		Name:     "lifthus_st",
		Value:    stSigned,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		Domain:   lifthus.CookieDomain,
		SameSite: http.SameSiteLaxMode,
	}
	c.SetCookie(cookie)

	// set pst to revoke it later.
	cookie2 := &http.Cookie{
		Name:     "lifthus_pst",
		Value:    stSigned,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		Domain:   lifthus.CookieDomain,
		Expires:  time.Now().AddDate(1, 0, 0),
		SameSite: http.SameSiteLaxMode,
	}
	c.SetCookie(cookie2)

	c.Response().Header().Set("Authorization", "Bearer "+stSigned)
	return c.String(http.StatusCreated, sid)
}

// HusSessionHandler doesn't need local version.
// it only communicates between servers.

// sessionSignHandler is a local development version of SessionSignHandler.
// which uses Authorization header instead of cookie.
func (ac authApiController) sessionSignHandler(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if !strings.HasPrefix(authorizationHeader, "Bearer ") || len(authorizationHeader) < 7 {
		return c.String(http.StatusBadRequest, "invalid authorization header")
	}
	lifthus_st := authorizationHeader[7:]

	// parse lifthus_st if it exists.
	lst, exp, err := helper.ParseJWTWithHMAC(lifthus_st)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusUnauthorized, err.Error())
	}

	sid := lst["sid"].(string)

	if exp {
		return c.String(http.StatusUnauthorized, "retry")
	}

	ls, err := db.QuerySessionBySID(c.Request().Context(), ac.dbClient, sid)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if ls == nil {
		err = fmt.Errorf("no valid session")
		log.Println(err)
		return c.String(http.StatusUnauthorized, err.Error())
	}

	// if it is already used to sign, return error.
	if ls.Used {
		err = fmt.Errorf("alredy used to sign")
		log.Println(err)
		return c.String(http.StatusUnauthorized, err.Error())
	}
	sidUUID, err := uuid.Parse(sid)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	// if the session is not used, set the session to be used.
	ac.dbClient.Session.UpdateOneID(sidUUID).SetUsed(true).Exec(c.Request().Context())

	// if the session is signed more than 5 seconds ago, revoke the session.
	if time.Since(*ls.SignedAt).Seconds() > 5 {
		_ = session.RevokeSession(c.Request().Context(), ac.dbClient, sid)
		err = fmt.Errorf("signing time is expired")
		log.Println(err)
		return c.String(http.StatusUnauthorized, err.Error())
	}

	nst := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"purpose": "lifthus_session",
		"sid":     sid,
		"uid":     strconv.FormatUint(*ls.UID, 10),
		"exp":     time.Now().Add(time.Minute * 5).Unix(),
	})
	nstSigned, err := nst.SignedString([]byte(lifthus.HusSecretKey))
	if err != nil {
		err = fmt.Errorf("signing accessToekn failed:%w", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set("Authorization", "Bearer "+nstSigned)

	// get user's Name from database using ls.UID
	lsu, err := db.QueryUserByUID(c.Request().Context(), ac.dbClient, *ls.UID)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// make struct with UID and Name
	signResp := struct {
		UID  string `json:"uid"`
		Name string `json:"username"`
	}{
		UID:  strconv.FormatUint(*ls.UID, 10),
		Name: lsu.Name,
	}
	signRespJSON, err := json.Marshal(signResp)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSONBlob(http.StatusOK, signRespJSON)
}
