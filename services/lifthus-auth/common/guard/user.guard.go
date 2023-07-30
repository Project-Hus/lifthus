package guard

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// UserGuard check if the client is signed.
// if there is embedded uid in context, it calls next handler,
// while it returns 401 when there is no uid in context.
func UserGuard(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		uidIntf := c.Get("uid")
		if uidIntf == nil {
			c.Response().Header().Set("WWW-Authenticate", `Bearer realm="auth.lifthus.com", error="not_signed"`)
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}
		_, ok := uidIntf.(uint64)
		if !ok {
			c.Response().Header().Set("WWW-Authenticate", `Bearer realm="auth.lifthus.com", error="not_signed"`)
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}

		return next(c)
	}
}
