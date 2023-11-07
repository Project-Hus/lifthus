package programqry

import (
	"log"
	"net/http"
	"routine/internal/repository"

	"github.com/labstack/echo/v4"
)

func SetProgramQueryControllerTo(e *echo.Echo) *echo.Echo {
	pc := &programQueryController{svc: newProgramQueryService()}
	e.GET("/routine/program/:code", pc.findProgramByCode)
	e.GET("/routine/programs", pc.findProgramsByTitle)
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

// findProgramsByTitle godoc
// @Router /programs [get]
// @Param title query string false "program title"
// @Summary
// @Tags program
// Success 200 "returns queried Programs"
// Failure 400 "invalid request"
// Failure 404 "not found"
// Failure 500 "failed to query Programs"
func (pc *programQueryController) findProgramsByTitle(c echo.Context) error {
	title := c.QueryParam("title")
	if title == "" {
		return c.String(http.StatusBadRequest, "invalid request")
	}
	qpDto, err := pc.svc.findProgramsByTitle(c.Request().Context(), title)
	if err != nil {
		log.Printf("failed to query Programs by title: %v", err)
		return c.String(http.StatusInternalServerError, "failed to query Programs")
	}
	return c.JSON(http.StatusOK, qpDto)
}
