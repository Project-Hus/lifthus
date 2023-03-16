package auth

import (
	"fmt"
	"lifthus-auth/common"
	"lifthus-auth/db"
	"lifthus-auth/helper"
	"lifthus-auth/service/session"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// NewSessionHandler godoc
// @Router       /session/new [post]
// @Summary      gets new connection and assign a lifthus session token.
// @Description  at the same time user connects to lifthus newly, the client requests new session token.
// @Description  and the server returns session id with session token in cookie.
// @ Description then the client send the session id to Hus auth server.
// @Tags         auth
// @Success      200 "session already exists"
// @Success      201 "returns session id with session token in cookie"
// @Failure      500 "failed to create new session"
func (ac authApiController) NewSessionHandler(c echo.Context) error {
	// get lifthus_st from cookie
	lifthus_st, err := c.Cookie("lifthus_st")
	if err != nil && err != http.ErrNoCookie {
		err = fmt.Errorf("[F]getting lifthus_st from cookie failed:%w", err)
		log.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	// if the client has token already, revoke it.
	if lifthus_st != nil {
		// parse lifthus_st
		st, exp, err := helper.ParseJWTwithHMAC(lifthus_st.Value)
		if err != nil || exp {
			err = fmt.Errorf("[F]parsing lifthus_st failed:%w", err)
			log.Println(err)
			return c.String(http.StatusUnauthorized, err.Error())
		}
		sid := st["sid"].(string)

		err = session.RevokeSession(c.Request().Context(), ac.Client, sid)
		if err != nil {
			err = fmt.Errorf("[F]revoking session failed:%w", err)
			log.Println(err)
			return c.String(http.StatusInternalServerError, err.Error())
		}
	}

	// create new lifthus session
	sid, stSigned, err := session.CreateSession(c.Request().Context(), ac.Client)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
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

	return c.String(http.StatusCreated, sid)
}

// HusSessionCheckHandler godoc
// @Router       /hus/session/check [post]
// @Summary      gets lifthus sid and uid from hus and set the login session.
// @Description  at the same time user connects to lifthus newly, the client requests new session token.
// @Description  and the server returns session id with session token in cookie.
// @Description then the client send the session id to Hus auth server.
// @Description and Hus validates the login session and tell lifthus.
// @Description and finally, Hus redirects the client to lifthus's endpoint.
// @Tags         auth
// @Success      200 "session checking success"
// @Failure      500 "failed to set the login session"
func (ac authApiController) HusSessionCheckHandler(c echo.Context) error {
	// from request body json, get sid string and uid string
	scbd := common.HusSessionCheckBody{}
	if err := c.Bind(&scbd); err != nil {
		log.Println("[F] binding HusSessionCheckBody failed: ", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	// Query if the user exists in the database
	u, err := db.QueryUserByUID(c.Request().Context(), ac.Client, scbd.Uid)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	if u == nil {
		// create new user if the user does not exist.
		_, err = db.CreateNewLifthusUser(c.Request().Context(), ac.Client, scbd)
		if err != nil {
			log.Println("[F] creating new user failed: ", err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	err = session.SetSignedSession(c.Request().Context(), ac.Client, scbd.Sid, scbd.Uid)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

// SessionCheckHandler godoc
// @Router       /session/access/newsid [post]
// @Summary      gets lifthus sid in cookie from client and set refresh token in cookie.
// @Description  Hus told lifthus that the user is logged in. so now we can set the login session.
// @Tags         auth
// @Success      200 "publishing refresh token success"
// @Failure      401 "unauthorized"
// @Failure      500 "internal server error"
func (ac authApiController) SessionCheckHandler(c echo.Context) error {
	// 세션 토큰 확인하고 로그인된 토큰이면 다시 허스에게 로그인 상태 확인하고
	//  세션 토큰은 재발급하면서 액세스 토큰 새로 발급
	fmt.Println("WOW you came here!")
	return c.NoContent(http.StatusOK)
}

// SessionCheckHandler godoc
// @Router       /session/access/newsid [post]
// @Summary      gets lifthus sid in cookie from client and set refresh token in cookie.
// @Description  Hus told lifthus that the user is logged in. so now we can set the login session.
// @Tags         auth
// @Success      200 "publishing refresh token success"
// @Failure      401 "unauthorized"
// @Failure      500 "internal server error"
func (ac authApiController) AccessHandler(c echo.Context) error {
	// 세션 토큰 확인하고 로그인된 토큰이면 다시 허스에게 로그인 상태 확인하고
	//  세션 토큰은 재발급하면서 액세스 토큰 새로 발급
	fmt.Println("WOW you came here!")
	return c.NoContent(http.StatusOK)
}
