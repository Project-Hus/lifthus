package auth

import (
	"fmt"
	"lifthus-auth/common"
	"lifthus-auth/db"
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
	if err != nil {
		if err != http.ErrNoCookie {
			log.Println("[F] getting lifthus_st from cookie failed: ", err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}
	// if there is session cookie already, just maintain the session.
	// it makes the session to be maintained through the whole tabs of the browser unlike using session storage.
	if lifthus_st != nil {
		// parse lifthus_st
		st, err := jwt.Parse(lifthus_st.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("HUS_SECRET_KEY")), nil
		})
		if err != nil || !st.Valid {
			return c.NoContent(http.StatusInternalServerError)
		}
		sid := st.Claims.(jwt.MapClaims)["sid"].(string)
		return c.String(http.StatusOK, sid)
	}
	/*
		although unlikely, this may cause malfunction. if the session is deleted from database before cookie,
		the client should clear the cookie self.
	*/

	// and if there is no cookie, create new session and return session id

	// create new session and get the session id
	sid, err := session.CreateSession(c.Request().Context(), ac.Client)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	// create new jwt session token with session id
	st := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sid": sid,
		"uid": nil, // it will be omitted actually.
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	hsk := []byte(os.Getenv("HUS_SECRET_KEY"))
	stSigned, err := st.SignedString(hsk)
	if err != nil {
		log.Println("[F] signing session token failed: ", err)
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
	//return c.Redirect(http.StatusPermanentRedirect, os.Getenv("HUS_AUTH_URL")+"/session/check/lifthus/"+sid)
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
// @Router       /session/check [post]
// @Summary      gets lifthus sid in cookie from client and set refresh token in cookie.
// @Description  Hus told lifthus that the user is logged in. so now we can set the login session.
// @Tags         auth
// @Success      200 "publishing refresh token success"
// @Failure      401 "unauthorized"
// @Failure      500 "internal server error"
func (ac authApiController) SessionCheckHandler(c echo.Context) error {
	fmt.Println("WOW you came here!")
	return c.NoContent(http.StatusOK)
}
