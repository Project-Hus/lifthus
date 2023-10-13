package actQuery

import (
	"log"
	"net/http"
	"routine/internal/repository"

	"github.com/labstack/echo/v4"
)

func SetActQueryControllerTo(e *echo.Echo) *echo.Echo {
	ac := &actQueryController{svc: newActQueryService()}
	e.GET("/routine/act", ac.queryAct)
	e.GET("/routine/acts", ac.queryActs)
	return e
}

type actQueryController struct {
	svc *actQueryService
}

// queryAct godoc
// @Router       /act [get]
// @Param code query string false "act code"
// @Summary      get specific act
// @Description
// @Tags         act
// @Success      200 "returns act info as json"
// @Failure      400 "invalid request"
// @Failure      404 "act not found"
// @Failure      500 "internal server error"
func (ac *actQueryController) queryAct(c echo.Context) error {
	code := c.QueryParam("code")
	qaDto, err := ac.svc.queryActByCode(c.Request().Context(), code)
	if repository.IsNotFound(err) {
		return c.String(http.StatusNotFound, "act not found")
	} else if err != nil {
		log.Printf("failed to query act: %v", err)
		return c.String(http.StatusInternalServerError, "failed to query act")
	}
	return c.JSON(http.StatusCreated, qaDto)
}

// queryActs godoc
// @Router       /acts [get]
// @Param name query string false "act code"
// @Summary      get acts that match the query
// @Description
// @Tags         act
// @Success      200 "returns acts info as json array"
// @Failure      400 "invalid request"
// @Failure      404 "act not found"
// @Failure      500 "internal server error"
func (ac *actQueryController) queryActs(c echo.Context) error {

	return c.String(http.StatusNotImplemented, "not implemented")
}
