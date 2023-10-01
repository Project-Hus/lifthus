package act

import (
	"lifthus-auth/common/guard"
	"net/http"
	"routine/internal/web/aws"
	"routine/internal/web/dto"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func SetActCommandControllerTo(e *echo.Echo) *echo.Echo {
	ac := &actController{svc: &actService{}, rb: &aws.RoutineBucket{}}
	e.POST("/routine/act", ac.createAct, guard.UserGuard)
	e.PUT("/routine/act", ac.upgradeAct, guard.UserGuard)
	return e
}

type actController struct {
	svc *actService
	rb  *aws.RoutineBucket
}

// createAct godoc
// @Router /act [post]
// @Param
// @Summary
// @Tags
// Success 201 "returns created Act"
// Failure 400 "invalid request"
// Failure 401 "unauthorized"
// Failure 403 "forbidden"
// Failure 500 "failed to create Act"
func (ac *actController) createAct(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		log.Errorf("failed to get multipart form: %v", err)
		return c.String(http.StatusInternalServerError, "failed to get multipart form")
	}
	files := form.File["images"]
	locations, err := ac.rb.UploadMultipartFilesToRoutineS3(files)
	if err != nil {
		log.Printf("failed to upload images to routine s3: %v", err)
	}
	dto := dto.CreateActDto{
		ActType:   c.FormValue("actType"),
		Name:      c.FormValue("name,"),
		Author:    c.FormValue("author"),
		Text:      c.FormValue("text"),
		ImageSrcs: locations,
	}

	act, err := ac.svc.createAct(dto)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to create Act")
	}
	return c.JSON(http.StatusCreated, act)
}

func (ac *actController) upgradeAct(c echo.Context) error {

	return nil
}
