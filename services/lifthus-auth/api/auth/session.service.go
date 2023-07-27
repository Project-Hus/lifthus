package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

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
//
// @Success      200 "Ok, session refreshed, session info JSON returned"
// @Success      201 "Created, new session issued, redirect to cloudhus and do connect"
// @Failure      500 "Internal Server Error"
func (ac authApiController) SessionHandler(c echo.Context) error {
	/*
		1. get sessoin token from cookie
		maybe cookie is not set or cookie is empty string. or maybe invalid
	*/
	lst, err := c.Cookie("lifthus_st")
	if err != nil && err != http.ErrNoCookie {
		return c.String(http.StatusInternalServerError, "failed to get cookie")
	}
	var rawLst string

	if lst != nil {
		rawLst = lst.Value
	}
	/*
		2. validate the session
	*/
	ls, err := session.ValidateSessionQueryUser(c.Request().Context(), rawLst)

	var nlst string // new session token

	/*
		3-1. try refresh the session
		if valid or expired but valid
	*/
	if err == nil || session.IsExpiredValid(err) {
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

		// returning sessoin user info
		var uinf *dto.SessionUserInfo

		// if session is signed by any user, return the user info
		if ls.Edges.User != nil {
			lsu := ls.Edges.User
			uinf = &dto.SessionUserInfo{
				UID:        strconv.FormatUint(lsu.ID, 10),
				Registered: lsu.Registered,
				Username:   lsu.Username,
				Usercode:   lsu.Usercode,
			}
		}

		// the client will get OK sign and that is all. no more thing to do.
		return c.JSON(http.StatusOK, uinf)
	}

	/*
		3-2. issue new session.
		first, after validation above, the session may turn out to be invalid.
		second, the refresh may have failed.
		in both cases, err won't be nil and the flow comes here.
		then issue a new session.
	*/
	ns, nlst, err := session.CreateSession(c.Request().Context())
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

	// the client will get Created sign.
	// in this case, the client must redirect to Cloudhus themselves to connect the sessions.
	return c.String(http.StatusCreated, ns.ID.String())
}

// GetSIDHandler godoc
// @Tags         auth
// @Router       /sid [get]
// @Summary		 returns client's SID. should be encrypted later.
//
// @Success      200 "Ok, session ID"
// @Failure      401 "Unauthorized, the token is expired"
// @Failure      500 "Internal Server Error"
func (ac authApiController) GetSIDHandler(c echo.Context) error {
	lst, err := c.Cookie("lifthus_st")
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to get cookie")
	}

	sid, _, exp, err := session.ValidateSession(c.Request().Context(), lst.Value)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to validate session")
	} else if exp {
		return c.String(http.StatusUnauthorized, "expired token")
	}

	return c.String(http.StatusOK, sid.String())
}

// SignInPropagationHandler godoc
// @Tags         auth
// @Router       /hus/signin [patch]
// @Summary		 processes user sign-in propagation from cloudhus.
// @Description  the "signin_propagation" token should be included in the request body.
// @Success      200 "Ok, session signed"
// @Failure      400 "Bad Request"
// @Failure	  500 "Internal Server Error"
func (ac authApiController) SignInPropagationHandler(c echo.Context) error {
	sipBytes, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to read body")
	}
	sip := string(sipBytes)

	// parse the token
	sipClaims, expired, err := helper.ParseJWTWithHMAC(sip)
	if expired || err != nil || sipClaims["pps"].(string) != "signin_propagation" {
		return c.String(http.StatusBadRequest, "invalid token")
	}

	// get the session ID
	sid := sipClaims["csid"].(string)
	suuid, err := uuid.Parse(sid)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid token")
	}
	hsid := sipClaims["hsid"].(string)
	hsuuid, err := uuid.Parse(hsid)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid token")
	}

	// query the session
	ls, err := db.QuerySessionByID(c.Request().Context(), suuid)
	if err != nil || ls == nil {
		return c.String(http.StatusInternalServerError, "failed to query session")
	}
	if ls.Hsid != nil && *ls.Hsid != hsuuid {
		return c.String(http.StatusBadRequest, "invalid token")
	}

	// marshal the "user" to json bytes
	cuB, err := json.Marshal(sipClaims["user"].(map[string]interface{}))
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid token")
	}
	// unmarshal it to HusConnUser
	var cu *dto.HusConnUser
	err = json.Unmarshal(cuB, &cu)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid token")
	}

	// query the user and register if not exists
	lu, err := db.QueryUserByID(c.Request().Context(), cu.Uid)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to query user")
	}
	if lu == nil {
		_, err := db.RegisterUser(c.Request().Context(), *cu)
		if err != nil {
			return c.String(http.StatusInternalServerError, "failed to register user")
		}
	}
	/* REFRESH THE USER INFO WITH THE LATEST ONE (not implemented yet) */

	_, err = ls.Update().SetUID(cu.Uid).SetSignedAt(time.Now()).Save(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to update session")
	}

	return c.String(http.StatusOK, "signed")
}

// SignOutPropagationHandler godoc
// @Tags         auth
// @Router       /hus/signout [patch]
// @Summary		 processes user sign-out propagation from cloudhus.
// @Description  the "signout_propagation" token should be included in the request body.
// @Success      200 "Ok, session signed"
// @Failure      400 "Bad Request"
// @Failure	  500 "Internal Server Error"
func (ac authApiController) SignOutPropagationHandler(c echo.Context) error {
	sopBytes, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to read body")
	}
	sop := string(sopBytes)

	// parse the token
	sopClaims, expired, err := helper.ParseJWTWithHMAC(sop)
	if expired || err != nil || sopClaims["pps"].(string) != "signout_propagation" {
		return c.String(http.StatusBadRequest, "invalid token")
	}

	hsid := sopClaims["hsid"].(string)
	hsuuid, err := uuid.Parse(hsid)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid token")
	}

	// query the latest related session
	ls, err := db.QuerySessionByHsid(c.Request().Context(), hsuuid)
	if err != nil || ls == nil {
		return c.String(http.StatusInternalServerError, "failed to query session")
	}

	// sign out of the session
	_, err = ls.Update().ClearUID().ClearSignedAt().Save(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to update session")
	}

	return c.String(http.StatusOK, "signed out")
}

// SignOutHandler godoc
// @Tags         auth
// @Router       /session/signout [patch]
// @Summary		 gets sign-out request from the client and propagates it to Cloudhus.
// @Success      200 "Ok, signed out of the session"
// @Failure      400 "Bad Request"
// @Failure	  401 "Unauthorized, the token is expired or the session is not signed"
// @Failure	  500 "Internal Server Error"
func (ac authApiController) SignOutHandler(c echo.Context) error {
	// get lifthus_st from the cookie
	lstSigned, err := c.Cookie("lifthus_st")
	if err != nil {
		log.Println("failed to get cookie")
		return c.String(http.StatusBadRequest, "failed to get token")
	}

	ls, err := session.ValidateSessionQueryUser(c.Request().Context(), lstSigned.Value)
	if session.IsExpiredValid(err) {
		c.Response().Header().Set("WWW-Authenticate", `Bearer realm="lifthus", error="expired token", error_description="the token is expired, refresh it`)
		log.Println("session expired")
		return c.String(http.StatusUnauthorized, "expired token")
	} else if err != nil {
		log.Println("failed to validate session")
		return c.String(http.StatusInternalServerError, "failed to validate session")
	} else if ls.Edges.User == nil {
		log.Println("the session is not signed")
		return c.String(http.StatusUnauthorized, "the session is not signed")
	}

	propagCh := make(chan error) // for waiting propagation goroutine
	txCh := make(chan error)     // for waiting transaction goroutine

	go func() {
		// generate new hus signout token
		sot := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"pps":  "hus_signout",
			"hsid": ls.Hsid.String(),
			"type": "total",
		})
		sotSigned, err := sot.SignedString(lifthus.HusSecretKeyBytes)
		if err != nil {
			propagCh <- fmt.Errorf("failed to sign token")
			return
		}
		// request to the Cloudhus endpoint
		req, err := http.NewRequest(http.MethodPatch, "https://auth.cloudhus.com/auth/hus/signout"+"", strings.NewReader(sotSigned))
		if err != nil {
			propagCh <- fmt.Errorf("failed to create request")
			return
		}
		resp, err := lifthus.Http.Do(req)
		if err != nil {
			propagCh <- fmt.Errorf("propagation failed")
			return
		}
		defer resp.Body.Close()
		// propagation failed or succeeded
		if resp.StatusCode == http.StatusOK {
			propagCh <- nil
			return
		}
		propagCh <- fmt.Errorf("propagation failed")
	}()

	tx, err := db.Client.Tx(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to start transaction")
	}

	go func() {
		ls, err = tx.Session.UpdateOne(ls).ClearUID().ClearSignedAt().SetTid(uuid.New()).Save(c.Request().Context())
		if err != nil {
			db.Rollback(tx, err)
			txCh <- err
			return
		}
		txCh <- nil
	}()

	propagErr := <-propagCh
	txErr := <-txCh

	if propagErr != nil || txErr != nil {
		_ = db.Rollback(tx, nil)
		log.Println("failed to sign out propag or tx failed")
		return c.String(http.StatusInternalServerError, "failed to sign out")
	}

	err = tx.Commit()
	if err != nil {
		_ = db.Rollback(tx, err)
		log.Println("failed to commit tx")
		return c.String(http.StatusInternalServerError, "failed to sign out")
	}

	lst, err := session.SessionToToken(c.Request().Context(), ls)
	if err != nil {
		log.Println("failed to tokenize the session")
		return c.String(http.StatusInternalServerError, "failed to sign out")
	}

	cookie := &http.Cookie{
		Name:     "lifthus_st",
		Value:    lst,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	c.SetCookie(cookie)

	log.Printf("user %d signed out", ls.Edges.User.ID)
	return c.String(http.StatusOK, "signed out")
}
