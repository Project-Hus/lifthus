package guard

import "github.com/labstack/echo/v4"

// UserGuard check if the client is signed.
// if there is embedded uid in context, it calls next handler,
// while it returns 401 when there is no uid in context.
func UserGuard(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		uidIntf := c.Get("uid")
		if uidIntf == nil {
			return c.String(401, "Unauthorized")
		}
		_, ok := uidIntf.(uint64)
		if !ok {
			return c.String(401, "Unauthorized")
		}

		return next(c)
	}
}
