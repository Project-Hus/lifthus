package guards

import "github.com/labstack/echo/v4"

func UserGuard(c func(c echo.Context) error) func(c echo.Context) error {
	return func(c echo.Context) error {
		return nil
	}
}
