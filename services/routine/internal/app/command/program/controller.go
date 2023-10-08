package programcmd

import (
	"lifthus-auth/common/guard"

	"github.com/labstack/echo/v4"
)

func SetProgramCommandControllerTo(e *echo.Echo) *echo.Echo {
	_ = &programCommandController{svc: newProgramCommandService()}
	e.POST("/routine/act", nil, guard.UserGuard)
	return e
}

type programCommandController struct {
	svc *programCommandService
}

// createAct godoc
// @Router /act [post]
// @Param Authorization header string true "lifthus_st"
// @Param creatActDto body dto.CreateActRequestDto true "create act dto"
// @Summary
// @Tags
// Success 201 "returns created Act"
// Failure 400 "invalid request"
// Failure 401 "unauthorized"
// Failure 403 "forbidden"
// Failure 500 "failed to create Act"
func (pc *programCommandController) createAct(c echo.Context) error {
	return nil
}
