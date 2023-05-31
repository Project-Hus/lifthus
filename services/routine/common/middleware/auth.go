package middleware

import (
	"net/http"
	"strings"

	"routine/common/helper"
	"routine/ent"

	"github.com/labstack/echo/v4"
)

func UidSetter(dbClient *ent.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var lifthus_st string

			org := c.Request().Header.Get("Origin")
			// local authentication
			if org == "http://localhost:3000" {
				// get authorization from req
				authHeader := c.Request().Header.Get("Authorization")
				if !strings.HasPrefix(authHeader, "Bearer ") {
					return next(c)
				}
				lifthus_st = authHeader[7:]
			} else { // production authentication
				authCookie, err := c.Request().Cookie("lifthus_st")
				if err != nil {
					if err == http.ErrNoCookie {
						return next(c)
					}
					return c.String(http.StatusInternalServerError, err.Error())
				}
				lifthus_st = authCookie.Value
			}
			uid, err := helper.ValidateLST(lifthus_st)
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}
			if uid != nil {
				c.Set("uid", uid)
			}
			return next(c)
		}
	}
}
