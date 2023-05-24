package guard

import "github.com/labstack/echo/v4"

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
