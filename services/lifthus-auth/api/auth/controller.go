package auth

import (
	"lifthus-auth/ent"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthApiControllerParams struct {
	DbClient   *ent.Client
	HttpClient *http.Client
}

// NewAuthApiController returns Echo comprising of auth api routes. instance to main.
func NewAuthApiController(authApi *echo.Echo, params AuthApiControllerParams) *echo.Echo {
	authApiController := newAuthApiController(params)

	authApi.GET("/auth", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Lifthus")
	})

	authApi.GET("/auth/session/new", authApiController.NewSessionHandler)
	authApi.PATCH("/auth/hus/session/sign", authApiController.HusSessionHandler)
	authApi.GET("/auth/session/sign", authApiController.SessionSignHandler)

	// authApi.DELETE("/auth/session/revoke", authApiController.SessionRevokeHandler)

	authApi.GET("/auth/session", authApiController.SessionHandler)
	authApi.PATCH("/auth/session/signout", authApiController.SignOutHandler)
	authApi.PATCH("/auth/hus/signin", authApiController.SignInPropagationHandler)
	authApi.PATCH("/auth/hus/signout", authApiController.SignOutPropagationHandler)

	return authApi
}

// newAuthApiController returns a new authApiController that implements every auth api features.
func newAuthApiController(params AuthApiControllerParams) authApis {
	return &authApiController{dbClient: params.DbClient, httpClient: params.HttpClient}
}

// authApiController defines what auth api has to have and implements authApis interface at service file.
type authApiController struct {
	dbClient   *ent.Client
	httpClient *http.Client
}

// authApis interface defines what auth api has to handle
type authApis interface {
	/* Session establishing process */
	// client -> lifthus -> client -> hus -> lifthus -> hus -> client -> SessionCheckHandler
	NewSessionHandler(c echo.Context) error  // from client
	HusSessionHandler(c echo.Context) error  // from hus
	SessionSignHandler(c echo.Context) error // from client

	SessionRevokeHandler(c echo.Context) error // from client

	SessionHandler(c echo.Context) error // from client
	SignOutHandler(c echo.Context) error // from client
	SignInPropagationHandler(c echo.Context) error
	SignOutPropagationHandler(c echo.Context) error
}
