package programqry

import (
	"log"
	"routine/internal/repository"

	"github.com/labstack/echo/v4"
)

func SetProgramQueryControllerTo(e *echo.Echo) *echo.Echo {
	pc := &programQueryController{svc: newProgramQueryService()}
	e.GET("/routine/program/:code", pc.findProgramByCode)
	return e
}

type programQueryController struct {
	svc *programQueryService
}

// findProgramByCode godoc
// @Router /program/{code} [get]
// @Param code path string true "program code"
// @Summary
// @Tags program
// Success 200 "returns queried Program"
// Failure 400 "invalid request"
// Failure 404 "not found"
// Failure 500 "failed to query Program"
func (pc *programQueryController) findProgramByCode(c echo.Context) error {
	code := c.Param("code")
	qpDto, err := pc.svc.findProgramByCode(c.Request().Context(), code)
	if repository.IsNotFound(err) {
		return c.String(404, "not found")
	} else if err != nil {
		log.Printf("failed to query Program by code: %v", err)
		return c.String(500, "failed to query Program")
	}
	return c.JSON(200, qpDto)
}
