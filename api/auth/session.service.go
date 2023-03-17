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
	// this handler would not be called if the client has unexpired access token.

	lifthus_psid, err := c.Cookie("lifthus_psid")
	if err != nil && err != http.ErrNoCookie {
		err = fmt.Errorf("!!getting lifthus_psid from cookie failed:%w", err)
		log.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if lifthus_psid != nil && lifthus_psid.Value != "" {
		err = session.RevokeSession(c.Request().Context(), ac.Client, lifthus_psid.Value)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, err.Error())
		}
	}

	// create new lifthus session
	sid, stSigned, err := session.CreateSession(c.Request().Context(), ac.Client)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// set cookie with session id
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

// AccessTokenHandler godoc
// @Router       /session/access [post]
// @Summary      gets lifthus sid in cookie from client and publishes access token.
// @Description  Hus told lifthus that the user is signed in.
// @Description so now we can publish access token to the client who has verified sid.
// @Description and also we revoke the used session token.
// @Tags         auth
// @Success      201 "publishing access token success"
// @Failure      401 "unauthorized"
// @Failure      500 "internal server error"
func (ac authApiController) AccessTokenHandler(c echo.Context) error {
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
	uid := lst["uid"].(string)

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

	err = session.RevokeSession(c.Request().Context(), ac.Client, sid)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if exp {
		err = fmt.Errorf("!!token is expired")
		log.Println(err)
		return c.String(http.StatusUnauthorized, err.Error())
	}

	// if the session is signed more than 5 seconds ago, it is expired.
	if time.Since(*ls.SignedAt).Seconds() > 5 {
		err = fmt.Errorf("!!signed session is expired")
		log.Println(err)
		return c.String(http.StatusUnauthorized, err.Error())
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sid": sid,
		"uid": uid,
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})
	atSigned, err := accessToken.SignedString([]byte(os.Getenv("HUS_SECRET_KEY")))
	if err != nil {
		err = fmt.Errorf("!!signing accessToekn failed:%w", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	atCookie := &http.Cookie{
		Name:     "lifthus_at",
		Value:    atSigned,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		Domain:   os.Getenv("LIFTHUS_DOMAIN"),
		SameSite: http.SameSiteDefaultMode,
	}
	c.SetCookie(atCookie)

	return c.String(http.StatusOK, "new access token and old session revoked")
}
