package rec

import (
	"net/http"
	"routine/ent"

	"github.com/labstack/echo/v4"
)

type ProgramApiControllerParams struct {
	DbClient   *ent.Client
	HttpClient *http.Client
}

// NewProgramApiController returns Echo instance comprising of program api routes to main.
func NewProgramApiController(programApi *echo.Echo, params ProgramApiControllerParams) *echo.Echo {
	programApiController := newProgramApiController(params)

	// return specific program
	programApi.GET("/routine/program/:program-slug", func(c echo.Context) error {
		return nil
	})
	// return
	programApi.GET("/routine/programs/title/:title", func(c echo.Context) error {
		return nil
	})
	programApi.POST("/routine/program", programApiController.createProgram)

	return programApi
}

// newAuthApiController returns a new authApiController that implements every auth api features.
func newProgramApiController(params ProgramApiControllerParams) programApis {
	return &programApiController{dbClient: params.DbClient, httpClient: params.HttpClient}
}

// authApiController defines what auth api has to have and implements authApis interface at service file.
type programApiController struct {
	dbClient   *ent.Client
	httpClient *http.Client
}

// authApis interface defines what auth api has to handle
type programApis interface {
	createProgram(c echo.Context) error
}
