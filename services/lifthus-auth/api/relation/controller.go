package relation

import (
	"lifthus-auth/ent"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RelationApiControllerParams struct {
	DbClient   *ent.Client
	HttpClient *http.Client
}

// NewRelationApiController returns Echo comprising of auth api routes. instance to main.
func NewRelationApiController(relationApi *echo.Echo, params RelationApiControllerParams) *echo.Echo {
	_ = newRelationApiController(params)

	relationApi.GET("/auth/relation", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Follower!")
	})

	return relationApi
}

// newAuthApiController returns a new authApiController that implements every auth api features.
func newRelationApiController(params RelationApiControllerParams) relationApis {
	return &relationApiController{dbClient: params.DbClient, httpClient: params.HttpClient}
}

// authApiController defines what auth api has to have and implements authApis interface at service file.
type relationApiController struct {
	dbClient   *ent.Client
	httpClient *http.Client
}

// authApis interface defines what auth api has to handle
type relationApis interface {
	GetUserFollowings(c echo.Context) error
	GetUserFollowers(c echo.Context) error

	// Unfollowing can be done with Follow endpoint.
	FollowUser(c echo.Context) error
}
