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

	authApi.GET("/auth/session", authApiController.SessionHandler)
	authApi.GET("/auth/sid", authApiController.GetSIDHandler)
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
	SessionHandler(c echo.Context) error // from client
	GetSIDHandler(c echo.Context) error  // from client
	SignOutHandler(c echo.Context) error // from client
	SignInPropagationHandler(c echo.Context) error
	SignOutPropagationHandler(c echo.Context) error
}
