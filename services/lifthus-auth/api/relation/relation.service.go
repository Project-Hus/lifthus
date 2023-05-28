package relation

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (rc relationApiController) GetUserFollowings(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Follower!")
}
func (rc relationApiController) GetUserFollowers(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Follower!")
}

func (rc relationApiController) FollowUser(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Follower!")
}
