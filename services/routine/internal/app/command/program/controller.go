package programcmd

import (
	"lifthus-auth/common/guard"
	"log"
	"net/http"
	"routine/internal/app/dto"

	"github.com/labstack/echo/v4"
)

func SetProgramCommandControllerTo(e *echo.Echo) *echo.Echo {
	pc := &programCommandController{svc: newProgramCommandService()}
	e.POST("/routine/act", pc.createProgram, guard.UserGuard)
	return e
}

type programCommandController struct {
	svc *programCommandService
}

// createProgram godoc
// @Router /program [post]
// @Param Authorization header string true "lifthus_st"
// @Param creatProgramDto body dto.CreateProgramRequestDto true "create program dto"
// @Summary
// @Tags program
// Success 201 "returns created Program"
// Failure 400 "invalid request"
// Failure 401 "unauthorized"
// Failure 403 "forbidden"
// Failure 500 "failed to create Program"
func (pc *programCommandController) createProgram(c echo.Context) error {
	cpDto := dto.CreateProgramRequestDto{}
	if err := c.Bind(&cpDto); err != nil {
		log.Printf("failed to bind request body: %v", err)
		return c.String(http.StatusBadRequest, "invalid request body")
	}
	cpSvcDto, err := cpDto.ToServiceDto()
	if err != nil {
		log.Printf("failed to convert createProgramReqDto to service dto: %v", err)
		return c.String(http.StatusBadRequest, "invalid request body")
	}
	clientId := c.Get("uid").(uint64)
	if cpSvcDto.Author != clientId {
		log.Printf("User %d attempted to create Program illegally", clientId)
		return c.String(http.StatusForbidden, "illegal access")
	}
	qpDto, err := pc.svc.createProgram(*cpSvcDto)
	if err != nil {
		log.Printf("failed to create Program: %v", err)
		return c.String(http.StatusInternalServerError, "failed to create Program")
	}
	return c.JSON(http.StatusCreated, qpDto)
}
