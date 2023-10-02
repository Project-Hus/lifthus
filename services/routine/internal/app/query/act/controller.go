package act

import (
	"lifthus-auth/common/guard"

	"github.com/labstack/echo/v4"
)

func SetActCommandControllerTo(e *echo.Echo) *echo.Echo {
	ac := &ActController{}
	e.POST("/routine/act", ac.createAct, guard.UserGuard)
	e.PUT("/routine/act", ac.upgradeAct, guard.UserGuard)
	return e
}

type ActController struct {
}

func (ac *ActController) createAct(c echo.Context) error {
	return nil
}

func (ac *ActController) upgradeAct(c echo.Context) error {
	return nil
}
