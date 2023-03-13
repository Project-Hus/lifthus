package auth

import (
	"lifthus-auth/ent"

	"github.com/labstack/echo/v4"
)

// authApis interface defines what auth api has to handle
type authApis interface {
	/* client side api */
	NewSessionHandler(c echo.Context) error

	/* hus side api */
	HusSessionCheckHandler(c echo.Context) error
}

// authApiController defines what auth api has to have and implements authApis interface at service file.
type authApiController struct {
	Client *ent.Client
}

// NewAuthApiController returns Echo comprising of auth api routes. instance to main.
func NewAuthApiController(client *ent.Client) *echo.Echo {
	authApi := echo.New()

	authApiController := newAuthApiController(client)

	authApi.POST("/session/new", authApiController.NewSessionHandler)
	authApi.POST("/hus/session/check", authApiController.HusSessionCheckHandler)

	return authApi
}

// newAuthApiController returns a new authApiController that implements every auth api features.
func newAuthApiController(client *ent.Client) authApis {
	return &authApiController{Client: client}
}
