package user

import (
	"fmt"
	"lifthus-auth/db"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (uc userApiController) GetUserInfo(c echo.Context) error {
	uid, err := strconv.ParseInt(c.Param("uid"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user, err := db.QueryUserByUID(c.Request().Context(), uc.dbClient, uint64(uid))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if user == nil {
		return c.String(http.StatusNotFound, fmt.Sprintf("user %d not found", uid))
	}

	return nil
}
