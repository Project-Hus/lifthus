package auth

import (
	"lifthus-auth/ent"

	"github.com/labstack/echo/v4"
)

// authApis interface defines what auth api has to handle
type authApis interface {
	/* Session establishing process */
	// client -> lifthus -> client -> hus -> lifthus -> hus -> client -> SessionCheckHandler
	NewSessionHandler(c echo.Context) error  // from client
	HusSessionHandler(c echo.Context) error  // from hus
	SessionSignHandler(c echo.Context) error // from client
}

// authApiController defines what auth api has to have and implements authApis interface at service file.
type authApiController struct {
	Client *ent.Client
}

// NewAuthApiController returns Echo comprising of auth api routes. instance to main.
func NewAuthApiController(client *ent.Client) *echo.Echo {
	authApi := echo.New()

	authApiController := newAuthApiController(client)

	authApi.GET("/session/new", authApiController.NewSessionHandler)
	authApi.POST("/hus/session/sign", authApiController.HusSessionHandler)
	authApi.POST("/session/sign", authApiController.SessionSignHandler)

	return authApi
}

// newAuthApiController returns a new authApiController that implements every auth api features.
func newAuthApiController(client *ent.Client) authApis {
	return &authApiController{Client: client}
}
