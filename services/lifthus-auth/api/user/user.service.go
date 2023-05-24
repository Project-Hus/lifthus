package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (uc userApiController) GetUserInfo(c echo.Context) error {
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	fmt.Println(uid)

	return nil
}
