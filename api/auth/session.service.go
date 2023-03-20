package auth

import (
	"fmt"
	"io"
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
			err := session.RevokeSessionToken(c.Request().Context(), ac.Client, lifthus_pst.Value)
			if err != nil {
				log.Println(err)
				return c.String(http.StatusInternalServerError, err.Error())
			}
		}
		// create new session
		sid, stSigned, err = session.CreateSession(c.Request().Context(), ac.Client)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// case B,C: if session token exists, validate it.
	} else {
		// ValidateSessionToken validates the token for case B,C, and updates the db for case B-1.
		sid, uid, exp, err = session.ValidateSession(c.Request().Context(), ac.Client, lifthus_st.Value)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, err.Error())

			// case B-1,2: if it is expired, refresh the token using same SID but clear the UID.
		} else if exp {
			stSigned, err = session.RefreshSessionToken(c.Request().Context(), ac.Client, sid)
			if err != nil {
				log.Println(err)
				return c.String(http.StatusInternalServerError, err.Error())
			}

			// case C: if it is valid, just keep the token and return
		} else {
			if uid != "" { // case C-1
				return c.String(http.StatusOK, uid)
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
		Domain:   os.Getenv("LIFTHUS_DOMAIN"),
		SameSite: http.SameSiteDefaultMode,
	}
	c.SetCookie(cookie)

	// set pst to revoke it later.
	cookie2 := &http.Cookie{
		Name:     "lifthus_pst",
		Value:    stSigned,
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
	body, err := io.ReadAll(c.Request().Body)
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
		"exp": time.Now().Add(time.Minute * 5).Unix(),
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
// @Summary      gets lifthus sid in cookie from client and signs the lifthus token.
// @Description  Hus told lifthus that the user is signed in.
// @Description so now we can sign the token which is owned by the client who has verified sid.
// @Tags         auth
// @Success      200 "session successfully signed"
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
