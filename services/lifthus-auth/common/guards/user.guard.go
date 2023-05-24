package guards

import "github.com/labstack/echo/v4"

func UserGuard(f func(c echo.Context) error) func(c echo.Context) error {
	return f
}
