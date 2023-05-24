package user

import (
	"lifthus-auth/common/guards"
	"lifthus-auth/ent"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserApiControllerParams struct {
	DbClient   *ent.Client
	HttpClient *http.Client
}

// NewUserApiController returns Echo comprising of auth api routes. instance to main.
func NewUserApiController(userApi *echo.Echo, params UserApiControllerParams) *echo.Echo {
	userApiController := newUserApiController(params)

	userApi.GET("/auth/user/:uid", guards.UserGuard(userApiController.GetUserInfo))

	return userApi
}

// newAuthApiController returns a new authApiController that implements every auth api features.
func newUserApiController(params UserApiControllerParams) userApis {
	return &userApiController{dbClient: params.DbClient, httpClient: params.HttpClient}
}

// authApiController defines what auth api has to have and implements authApis interface at service file.
type userApiController struct {
	dbClient   *ent.Client
	httpClient *http.Client
}

// authApis interface defines what auth api has to handle
type userApis interface {
	GetUserInfo(c echo.Context) error
}
