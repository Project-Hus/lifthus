package actCommand

import (
	"lifthus-auth/common/guard"
	"net/http"
	"routine/internal/app/dto"

	"log"

	"github.com/labstack/echo/v4"
)

func SetActCommandControllerTo(e *echo.Echo) *echo.Echo {
	ac := &actController{svc: newActService()}
	e.POST("/routine/act", ac.createAct, guard.UserGuard)
	e.PUT("/routine/act", ac.upgradeAct, guard.UserGuard)
	return e
}

type actController struct {
	svc *actService
}

// createAct godoc
// @Router /act [post]
// @Param author formData string true "author of act"
// @Param name formData string true "name of act"
// @param actType formData string true "type of act"
// @Param text formData string true "text of act"
// @Param images formData file false "images of act"
// @Param Authorization header string true "lifthus_st"
// @Summary
// @Tags
// Success 201 "returns created Act"
// Failure 400 "invalid request"
// Failure 401 "unauthorized"
// Failure 403 "forbidden"
// Failure 500 "failed to create Act"
func (ac *actController) createAct(c echo.Context) error {
	locations, err := getMultipartFormAndUploadActImagesToRoutineS3(c)
	if err != nil {
		log.Printf("failed to upload images to routine s3: %v", err)
		return c.String(http.StatusInternalServerError, "failed to upload images to routine s3")
	}

	dto, err := parseFormAndGenerateCreateActDto(c, locations)
	if err != nil {
		log.Printf("constructing CreateActDto failed: %v", err)
		return c.String(http.StatusBadRequest, "failed to construct dto")
	}

	clientId := c.Get("uid").(uint64)
	if dto.Author != clientId {
		log.Printf("illegal access: %v", err)
		return c.String(http.StatusForbidden, "illegal access")
	}

	act, err := ac.svc.createAct(c.Request().Context(), *dto)
	if err != nil {
		log.Printf("failed to create Act: %v", err)
		return c.String(http.StatusInternalServerError, "failed to create Act")
	}
	return c.JSON(http.StatusCreated, act)
}

func (ac *actController) upgradeAct(c echo.Context) error {
	ac.svc.upgradeAct(c.Request().Context(), dto.UpgradeActDto{})
	return nil
}
