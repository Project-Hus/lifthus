package auth

import (
	"fmt"
	"io/ioutil"
	"lifthus-auth/common"
	"lifthus-auth/db"
	"lifthus-auth/helper"

	"lifthus-auth/service/session"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// NewSessionHandler godoc
// @Router       /session/new [post]
// @Summary      when lifthus web app is opened, session token is assigned.
// @Description  at the same time the user opens Lifthus from browser, the client requests new session token.
// @Description  and Lifthus auth server returns session id with session token in cookie.
// @ Description then the client send the session id to Hus auth server to validate the session.
// @Tags         auth
// @Success      201 "returns session id with session token in cookie"
// @Failure      500 "failed to create new session"
func (ac authApiController) NewSessionHandler(c echo.Context) error {
	/* 4 ways to handle session token */
	// A: if there is no session, return new session token with new SID
	// B-1: if it is signed but expired, reset used, signed_at, uid from db and return new token with same SID
	// B-2: if it is not signed and expired, return new session token with same SID
	// C: if it is valid, just return

	// first get psid and st from cookie
	lifthus_psid, err := c.Cookie("lifthus_psid")
	if err != nil && err != http.ErrNoCookie {
		err = fmt.Errorf("!!getting lifthus_psid from cookie failed:%w", err)
		log.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	lifthus_st, err := c.Cookie("lifthus_st")
	if err != nil && err != http.ErrNoCookie {
		log.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// sid and stSigned to be set in cookie
	var sid, stSigned string

	// case A: if there is no session, revoke psid and create new session.
	if lifthus_st == nil || lifthus_st.Value == "" {
		// revoke psid if psid exists
		if lifthus_psid != nil && lifthus_psid.Value != "" {
			err := session.RevokeSession(c.Request().Context(), ac.Client, lifthus_psid.Value)
			if err != nil {
				log.Println(err)
				return c.String(http.StatusInternalServerError, err.Error())
			}
			// create new session
			sid, stSigned, err = session.CreateSession(c.Request().Context(), ac.Client)
			if err != nil {
				log.Println(err)
				return c.String(http.StatusInternalServerError, err.Error())
			}
		}

		// case B,C: if session token exists, validate it.
	} else {
		// ValidateSessionToken validates the token for case B, and updates the db for case B-1.
		sid, uid, exp, err := session.ValidateSessionToken(c.Request().Context(), ac.Client, lifthus_st.Value)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, err.Error())

			// case B: if it is expired, refresh the token using same SID.
		} else if exp {
			stSigned, err = session.RefreshSession(c.Request().Context(), ac.Client, sid)
			if err != nil {
				log.Println(err)
				return c.String(http.StatusInternalServerError, err.Error())
			}
			// case C: if it is valid, just keep the token and return
		} else {
			return c.String(http.StatusOK, uid)

			// case B: if it is expired, refresh the token using same SID.
		}
	}

	// set session token in cookie.
	cookie := &http.Cookie{
		Name:     "lifthus_st",
		Value:    stSigned,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		Domain:   os.Getenv("LIFTHUS_DOMAIN"),
		SameSite: http.SameSiteDefaultMode,
	}
	c.SetCookie(cookie)

	// set psid to revoke it later.
	cookie2 := &http.Cookie{
		Name:     "lifthus_psid",
		Value:    sid,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		Domain:   os.Getenv("LIFTHUS_DOMAIN"),
		Expires:  time.Now().AddDate(1, 0, 0),
		SameSite: http.SameSiteDefaultMode,
	}
	c.SetCookie(cookie2)

	return c.String(http.StatusCreated, sid)
}

// HusSessionHandler godoc
// @Router       /hus/session/sign [post]
// @Summary      gets lifthus sid and uid from Hus and sets the session token to be signed in.
// @Description Hus sends SID and UID which are verified and Lifthus sets the session token to be signed in.
// @Tags         auth
// @Success      200 "session signing success"
// @Failure      500 "failed to set the login session"
func (ac authApiController) HusSessionHandler(c echo.Context) error {

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		err = fmt.Errorf("!!reading request body failed:%w", err)
		log.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	hscb := string(body)

	hscbParsed, exp, err := helper.ParseJWTwithHMAC(hscb)
	if err != nil {
		err = fmt.Errorf("!!parsing jwt failed:%w", err)
		log.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	} else if exp {
		err = fmt.Errorf("!!token is expired")
		log.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	hscbParsed["sid"] = hscbParsed["sid"].(string)
	// from request body json, get sid string and uid string
	scbd := common.HusSessionCheckBody{
		Sid:           hscbParsed["sid"].(string),
		Uid:           hscbParsed["uid"].(string),
		Email:         hscbParsed["email"].(string),
		EmailVerified: hscbParsed["email_verified"].(bool),
		Name:          hscbParsed["name"].(string),
		GivenName:     hscbParsed["given_name"].(string),
		FamilyName:    hscbParsed["family_name"].(string),
		Birthdate:     hscbParsed["birthdate"].(string),
	}

	// Query if the user exists in the database
	u, err := db.QueryUserByUID(c.Request().Context(), ac.Client, scbd.Uid)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if u == nil {
		// create new user if the user does not exist.
		_, err = db.CreateNewLifthusUser(c.Request().Context(), ac.Client, scbd)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, err.Error())
		}
	}

	err = session.SetSignedSession(c.Request().Context(), ac.Client, scbd.Sid, scbd.Uid)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	nst := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sid": scbd.Sid,
		"uid": scbd.Uid,
		"exp": time.Now().Add(time.Minute * 10).Unix(),
	})
	nstSigned, err := nst.SignedString([]byte(os.Getenv("HUS_SECRET_KEY")))
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	nstCookie := &http.Cookie{
		Name:     "lifthus_st",
		Value:    nstSigned,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		Domain:   os.Getenv("LIFTHUS_DOMAIN"),
		SameSite: http.SameSiteDefaultMode,
	}

	c.SetCookie(nstCookie)

	return c.String(http.StatusOK, "session signing success")
}

// SessionSignHandler godoc
// @Router       /session/sign [post]
// @Summary      gets lifthus sid in cookie from client and publishes access token.
// @Description  Hus told lifthus that the user is signed in.
// @Description so now we can publish access token to the client who has verified sid.
// @Description and also we revoke the used session token.
// @Tags         auth
// @Success      201 "publishing access token success"
// @Failure      401 "unauthorized"
// @Failure      500 "internal server error"
func (ac authApiController) SessionSignHandler(c echo.Context) error {
	// get lifthus_st from cookie
	lifthus_st, err := c.Cookie("lifthus_st")
	if err != nil {
		err = fmt.Errorf("!!no valid token:%w", err)
		return c.String(http.StatusUnauthorized, err.Error())
	}

	// parse lifthus_st
	lst, exp, err := helper.ParseJWTwithHMAC(lifthus_st.Value)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusUnauthorized, err.Error())
	}

	sid := lst["sid"].(string)

	if exp {
		err = session.RevokeSession(c.Request().Context(), ac.Client, sid)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, err.Error())
		}
		err = fmt.Errorf("!!token is expired")
		log.Println(err)
		return c.String(http.StatusUnauthorized, err.Error())
	}

	ls, err := db.QuerySessionBySID(c.Request().Context(), ac.Client, sid)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if ls == nil {
		err = fmt.Errorf("!!no valid session")
		log.Println(err)
		return c.String(http.StatusUnauthorized, err.Error())
	}

	// if the session is signed more than 5 seconds ago, it is expired.
	if time.Since(*ls.SignedAt).Seconds() > 5 {
		err = session.RevokeSession(c.Request().Context(), ac.Client, sid)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, err.Error())
		}
		err = fmt.Errorf("!!signed session is expired")
		log.Println(err)
		return c.String(http.StatusUnauthorized, err.Error())
	}

	nst := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sid": sid,
		"uid": ls.UID,
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})
	nstSigned, err := nst.SignedString([]byte(os.Getenv("HUS_SECRET_KEY")))
	if err != nil {
		err = fmt.Errorf("!!signing accessToekn failed:%w", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	nstCookie := &http.Cookie{
		Name:     "lifthus_st",
		Value:    nstSigned,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		Domain:   os.Getenv("LIFTHUS_DOMAIN"),
		SameSite: http.SameSiteDefaultMode,
	}
	c.SetCookie(nstCookie)

	return c.String(http.StatusOK, ls.UID.String())
}
