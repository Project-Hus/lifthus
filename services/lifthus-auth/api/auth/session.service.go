package auth

import (
	"encoding/json"
	"fmt"
	"io"

	"lifthus-auth/common/db"
	"lifthus-auth/common/dto"
	"lifthus-auth/common/helper"
	"lifthus-auth/common/lifthus"
	"strconv"

	"lifthus-auth/common/service/session"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// SessionHandler godoc
// @Tags         auth
// @Router       /session [get]
// @Summary		 validates session. publishes new one if it isn't. refreshes expired session.
// @Success      200 "Ok, session refreshed, session info JSON returned"
// @Success      201 "Created, new session issued, redirect to cloudhus and do connect"
// @Failure      500 "Internal Server Error"
func (ac authApiController) SessionHandler(c echo.Context) error {
	/*
		1. get session token from cookie
		  the token string may be empty if there's no cookie.
		  and may the token be invalid.
		2. validate the token
		  if expired, try refresh
		3-1. if session invalid or refresh failed, issue a new session
		3-2. if session valid,
		  try connect to Cloudhus if it still isn't connected.
		  if connected, just return it.
	*/
	// first get the Lifthus session token from cookie
	lst, err := c.Cookie("lifthus_st")
	if err != nil && err != http.ErrNoCookie {
		return c.String(http.StatusInternalServerError, "failed to get cookie")
	}
	// validate the session
	ls, err := session.ValidateSessionV2(c.Request().Context(), lst.Value)

	var nlst string // new session token

	// try refresh if valid or expired
	if err == nil || session.IsExpired(err) {
		ls, nlst, err = session.RefreshSessionHard(c.Request().Context(), ls)
	}

	// if refresh succeeded, return the refreshed session token
	if err == nil {
		nlstCookie := &http.Cookie{
			Name:     "lifthus_st",
			Value:    nlst,
			Path:     "/",
			Domain:   ".lifthus.com",
			HttpOnly: true,
			Secure:   lifthus.CookieSecure,
			SameSite: http.SameSiteLaxMode,
		}
		c.SetCookie(nlstCookie)

		return c.JSON(http.StatusOK, struct {
			UID *uint64 `json:"uid"`
		}{
			UID: ls.UID,
		})
	}

	// create new session in case the session is invalid or refresh failed
	_, nlst, err = session.CreateSessionV2(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to issue new session")
	}
	nlstCookie := &http.Cookie{
		Name:     "lifthus_st",
		Value:    nlst,
		Path:     "/",
		Domain:   ".lifthus.com",
		HttpOnly: true,
		Secure:   lifthus.CookieSecure,
		SameSite: http.SameSiteLaxMode,
	}
	c.SetCookie(nlstCookie)

	return c.String(http.StatusCreated, "new session issued")
}

// NewSessionHandler godoc
// @Router       /session/new [get]
// @Summary      accepts tokens in cookie, parse and validate them, and returns tokens depending on the token's status.
// @Description  case A: no session, return newly generated session token with 201.
// @Description  case B-1: signed but expired, reset session info(used, signed_at, uid) except SID and return new session token with 201.
// @Description case B-2: not signed and expired, return new session token keeping SID with 201.
// @Description case C-1: valid and signed, just return with 200.
// @Description case C-2: valid and not signed, return with 201 to tell client to check Hus session.
// @Tags         auth
// @Success      200 "if valid session exists, return uid"
// @Success      201 "if there's no session or existing session is expired, return new session token"
// @Failure      500 "failed to create new session"
func (ac authApiController) NewSessionHandler(c echo.Context) error {
	origin := c.Request().Header.Get("Origin")
	if origin == "http://localhost:3000" {
		return ac.newSessionHandler(c)
	}

	lifthus_pst, err := c.Cookie("lifthus_pst")
	if err != nil && err != http.ErrNoCookie {
		log.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	lifthus_st, err := c.Cookie("lifthus_st")
	if err != nil && err != http.ErrNoCookie {
		log.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var sid, stSigned, uid string
	var exp bool

	// case A: no session, create new session
	if lifthus_st == nil || lifthus_st.Value == "" {
		// revoke psid if psid exists
		if lifthus_pst != nil && lifthus_pst.Value != "" {
			err := session.RevokeSessionToken(c.Request().Context(), ac.dbClient, lifthus_pst.Value)
			if err != nil {
				log.Println(err)
				return c.String(http.StatusInternalServerError, err.Error())
			}
		}
		// create new session
		sid, stSigned, err = session.CreateSession(c.Request().Context(), ac.dbClient)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// case B,C: if session token exists, validate it.
	} else {
		// ValidateSessionToken validates the token for case B,C, and updates the db for case B-1.
		sid, uid, exp, err = session.ValidateSession(c.Request().Context(), ac.dbClient, lifthus_st.Value)
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
				return c.JSONBlob(http.StatusOK, keepRespJSON)
			} else { // case C-2
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

	return c.String(http.StatusCreated, sid)
}

// HusSessionHandler godoc
// @Router       /hus/session/sign [patch]
// @Summary      gets Hus id token and sets the session token to be signed in after updating the info.
// @Description Hus sends id token and Lifthus sets the session info to be signed in with specific uid.
// @Description and if the user is not registered, Lifthus will register the user.
// @Description Hus user info's change will be reflected as well.
// @Tags         auth
// @Success      200 "session signing success"
// @Failure      500 "failed to set the login session"
func (ac authApiController) HusSessionHandler(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		err = fmt.Errorf("reading request body failed:%w", err)
		log.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	hscb := string(body)

	hscbParsed, exp, err := helper.ParseJWTWithHMAC(hscb)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	} else if exp {
		err = fmt.Errorf("token is expired")
		log.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// from request body json, get sid string and uid string
	scbd := dto.HusSessionCheckBody{
		Sid:             hscbParsed["sid"].(string),
		Uid:             hscbParsed["uid"].(string),
		ProfileImageURL: hscbParsed["profile_image_url"].(string),
		Email:           hscbParsed["email"].(string),
		EmailVerified:   hscbParsed["email_verified"].(bool),
		Name:            hscbParsed["name"].(string),
		GivenName:       hscbParsed["given_name"].(string),
		FamilyName:      hscbParsed["family_name"].(string),
		Birthdate:       hscbParsed["birthdate"].(string),
	}

	// Query if the user exists in the database
	scbdUidUint64, err := strconv.ParseUint(scbd.Uid, 10, 64)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	u, err := db.QueryUserByUID(c.Request().Context(), ac.dbClient, scbdUidUint64)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if u == nil {
		// create new user if the user does not exist.
		_, err = db.CreateNewLifthusUser(c.Request().Context(), ac.dbClient, scbd)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, err.Error())
		}
	}

	err = session.SignSession(c.Request().Context(), ac.dbClient, scbd.Sid, scbdUidUint64)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "session signing success")
}

// SessionSignHandler godoc
// @Router       /session/sign [get]
// @Summary      gets lifthus sid in cookie from client, and signs the lifthus token.
// @Description  Hus told lifthus that the client with specific SID is signed in.
// @Description so now we can sign the token which is owned by the client who has verified sid.
// @Tags         auth
// @Success      200 "session successfully signed"
// @Failure      401 "unauthorized"
// @Failure      500 "internal server error"
func (ac authApiController) SessionSignHandler(c echo.Context) error {
	origin := c.Request().Header.Get("Origin")
	if origin == "http://localhost:3000" {
		return ac.sessionSignHandler(c)
	}

	// get lifthus_st from cookie
	lifthus_st, err := c.Cookie("lifthus_st")
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	// parse lifthus_st if it exists.
	lst, exp, err := helper.ParseJWTWithHMAC(lifthus_st.Value)
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
	nstCookie := &http.Cookie{
		Name:     "lifthus_st",
		Value:    nstSigned,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		Domain:   lifthus.CookieDomain,
		SameSite: http.SameSiteLaxMode,
	}
	c.SetCookie(nstCookie)

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

// SessionRevokeHandler godoc
// @Router       /session/revoke [delete]
// @Summary      revokes lifthus session
// @Tags         auth
// @Success      200 "if valid session exists, return uid"
// @Failure      404 "token not found"
// @Failure      500 "failed to create new session"
func (ac authApiController) SessionRevokeHandler(c echo.Context) error {
	lifthus_st, err := c.Cookie("lifthus_st")
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}
	// validate lifthus_st
	sid, _, _, err := session.ValidateSession(c.Request().Context(), ac.dbClient, lifthus_st.Value)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}
	sidUUID, err := uuid.Parse(sid)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	_ = ac.dbClient.Session.DeleteOneID(sidUUID).Exec(c.Request().Context())

	revokedCookie := &http.Cookie{
		Name:     "lifthus_st",
		Value:    "",
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		Domain:   lifthus.CookieDomain,
		SameSite: http.SameSiteLaxMode,
	}
	revokedCookie2 := &http.Cookie{
		Name:     "lifthus_pst",
		Value:    "",
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		Domain:   lifthus.CookieDomain,
		SameSite: http.SameSiteLaxMode,
	}
	c.SetCookie(revokedCookie)
	c.SetCookie(revokedCookie2)

	return c.String(http.StatusOK, "session revoked")
}
