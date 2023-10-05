package actQuery

import (
	"log"
	"net/http"
	"routine/internal/repository"

	"github.com/labstack/echo/v4"
)

func SetActQueryControllerTo(e *echo.Echo) *echo.Echo {
	ac := &actQueryController{svc: newActQueryService()}
	e.GET("/routine/act", ac.queryActByCode)
	return e
}

type actQueryController struct {
	svc *actQueryService
}

// queryActByCode godoc
// @Router       /act/{code} [get]
// @Param code path string true "act code"
// @Summary      get act by code
// @Description
// @Tags         act
// @Success      200 "returns act info as json"
// @Failure      400 "invalid request"
// @Failure      404 "act not found"
// @Failure      500 "internal server error"
func (ac *actQueryController) queryActByCode(c echo.Context) error {
	code := c.Param("code")
	qaDto, err := ac.svc.queryActByCode(c.Request().Context(), code)
	if repository.IsNotFound(err) {
		return c.String(http.StatusNotFound, "act not found")
	} else if err != nil {
		log.Printf("failed to query act: %v", err)
		return c.String(http.StatusInternalServerError, "failed to query act")
	}
	return c.JSON(http.StatusCreated, actDto)
}
