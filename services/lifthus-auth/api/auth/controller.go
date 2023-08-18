package auth

import (
	"context"
	"lifthus-auth/common/guard"
	"lifthus-auth/common/helper"
	"net/http"

	"log"

	"github.com/labstack/echo/v4"
)

// NewAuthApiController takes Echo instance and attaches auth API controller to it.
func NewAuthApiController(authApi *echo.Echo) *echo.Echo {

	authApi.GET("/auth", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Lifthus")
	})

	authApi.GET("/auth/session", authController.SessionHandler)
	authApi.GET("/auth/sid", authController.GetSIDHandler)
	authApi.PATCH("/auth/session/signout", authController.SignOutHandler, guard.UserGuard)
	authApi.PATCH("/auth/hus/signin", authController.SignInPropagationHandler)
	authApi.PATCH("/auth/hus/signout", authController.SignOutPropagationHandler)

	// new desgin patterns below
	authApi.DELETE("/auth/account", authController.DeleteAccountController, guard.UserGuard)

	return authApi
}

// init initializes auth controller and service
func init() {
	authService = newAuthService()
	authController = newAuthApiController()
}

// SessionHandler godoc
// @Tags         auth
// @Router       /auth/account [delete]
// @Summary		 deletes user's lifthus account
// @Success      200 "Ok, the account is deleted"
// @Success      400 "Bad Request, invalid request"
// @Failure      500 "Internal Server Error"
func (ac authApiController) DeleteAccountController(c echo.Context) error {
	uid, ok := c.Get("uid").(uint64)
	if !ok {
		return c.String(400, "invalid uid")
	}

	// get lifthus_st from the cookie
	lst, err := helper.GetHeaderLST(c)
	if err != nil {
		log.Println("failed to get lst")
		return c.String(http.StatusBadRequest, "no valid token")
	}

	nlst, err := authService.signOutService(c.Request().Context(), lst)
	if err != nil {
		log.Printf("failed to sign out: %v", err)
		return c.String(http.StatusInternalServerError, "failed to sign out")
	}

	c = helper.SetAuthHeader(c, nlst)

	log.Printf("user %d signed out", uid)

	err = authService.DeleteAccountService(c.Request().Context(), uid)
	if err != nil {
		log.Printf("failed to delete account: %v", err)
		return c.String(500, "failed to delete account")
	}

	return c.String(200, "account deleted")
}

/* ========== Controller declaration ========== */
// authController is a controller set that handles auth api requests.
var authController authApis

// authApiController is a controller struct that implements authApis interface.
type authApiController struct {
}

// newAuthApiController returns authApiController instance.
func newAuthApiController() authApis {
	return &authApiController{}
}

// authApis interface defines what auth api has to handle
type authApis interface {
	SessionHandler(c echo.Context) error // from client
	GetSIDHandler(c echo.Context) error  // from client
	SignOutHandler(c echo.Context) error // from client
	SignInPropagationHandler(c echo.Context) error
	SignOutPropagationHandler(c echo.Context) error

	// new design patterns below
	DeleteAccountController(c echo.Context) error
}

/* ========== Service declaration ========== */
// authService is a service set that handles auth api requests.
var authService authApiServices

// authApiService is a service struct that implements authApiServices interface.
type authApiService struct {
}

// newAuthService returns authApiService instance.
func newAuthService() authApiServices {
	return &authApiService{}
}

// authApiServices interface defines what services should be implemented
type authApiServices interface {
	DeleteAccountService(c context.Context, uid uint64) error

	signOutService(c context.Context, lst string) (nlst string, err error)
}
