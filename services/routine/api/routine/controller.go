package routine

import (
	"net/http"
	"routine/ent"

	"github.com/labstack/echo/v4"
)

type RoutineApiControllerParams struct {
	DbClient   *ent.Client
	HttpClient *http.Client
}

// NewRoutineApiController returns Echo instance comprising of routine api routes to main.
func NewRoutineApiController(routineApi *echo.Echo, params RoutineApiControllerParams) *echo.Echo {
	routineApiController := newRoutineApiController(params)

	routineApi.GET("/routine", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Lifthus")
	})

	routineApi.GET("/routine/tmp", routineApiController.tmp)

	return routineApi
}

// newAuthApiController returns a new authApiController that implements every auth api features.
func newRoutineApiController(params RoutineApiControllerParams) routineApis {
	return &routineApiController{dbClient: params.DbClient, httpClient: params.HttpClient}
}

// authApiController defines what auth api has to have and implements authApis interface at service file.
type routineApiController struct {
	dbClient   *ent.Client
	httpClient *http.Client
}

// authApis interface defines what auth api has to handle
type routineApis interface {
	tmp(c echo.Context) error
}
