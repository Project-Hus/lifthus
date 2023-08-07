package middleware

import (
	"lifthus-auth/common/helper"
	"lifthus-auth/common/service/session"

	"net/http"

	"github.com/labstack/echo/v4"
)

func UidSetter() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			var lst, err = helper.GetHeaderLST(c)
			if err != nil {
				return next(c)
			}

			_, uid, exp, err := session.ValidateSession(c.Request().Context(), lst)
			if err != nil {
				return c.String(http.StatusInternalServerError, "illegal session")
			}
			if uid != nil && !exp {
				c.Set("uid", *uid)
			}
			if exp {
				c.Set("exp", true)
			} else {
				c.Set("exp", false)
			}
			return next(c)
		}
	}
}
