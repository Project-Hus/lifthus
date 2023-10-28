package actcmd

import (
	"lifthus-auth/common/guard"
	"net/http"
	"routine/internal/app/dto"

	"log"

	"github.com/labstack/echo/v4"
)

func SetActCommandControllerTo(e *echo.Echo) *echo.Echo {
	ac := &actCommandController{svc: newActCommandService()}
	e.POST("/routine/act", ac.createAct, guard.UserGuard)
	return e
}

type actCommandController struct {
	svc *actCommandService
}

// createAct godoc
// @Router /act [post]
// @Param Authorization header string true "lifthus_st"
// @Param creatActDto body dto.CreateActRequestDto true "create act dto"
// @Summary
// @Tags act
// Success 201 "returns created Act"
// Failure 400 "invalid request"
// Failure 401 "unauthorized"
// Failure 403 "forbidden"
// Failure 500 "failed to create Act"
func (ac *actCommandController) createAct(c echo.Context) error {
	reqDto := dto.CreateActRequestDto{}
	if err := c.Bind(&reqDto); err != nil {
		log.Printf("failed to bind request: %v", err)
		return c.String(http.StatusBadRequest, "invalid request")
	}

	svcDto, err := reqDto.ToServiceDto()
	if err != nil {
		log.Printf("failed to convert to service dto: %v", err)
		return c.String(http.StatusBadRequest, "invalid request")
	}

	clientId := c.Get("uid").(uint64)
	if svcDto.Author != clientId {
		log.Printf("illegal access: %v", err)
		return c.String(http.StatusForbidden, "illegal access")
	}

	act, err := ac.svc.createAct(c.Request().Context(), *svcDto)
	if err != nil {
		log.Printf("failed to create Act: %v", err)
		return c.String(http.StatusInternalServerError, "failed to create Act")
	}
	return c.JSON(http.StatusCreated, act)
}
