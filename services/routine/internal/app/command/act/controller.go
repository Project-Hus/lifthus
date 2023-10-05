package actCommand

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
	e.POST("/routine/act/upgrade", ac.upgradeAct, guard.UserGuard)
	return e
}

type actCommandController struct {
	svc *actCommandService
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
func (ac *actCommandController) createAct(c echo.Context) error {
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

// upgradeAct godoc
// @Router /act/upgarde [post]
// @Param author formData string true "author of act"
// @Summary
// @Tags
// Success 201 "returns created Act"
// Failure 400 "invalid request"
// Failure 401 "unauthorized"
// Failure 403 "forbidden"
// Failure 500 "failed to create Act"
func (ac *actCommandController) upgradeAct(c echo.Context) error {
	// uaDto := dto.UpgradeActDto{
	// 	ActCode: c.FormValue("actCode"),
	// 	Text:    c.FormValue("text"),
	// }
	// //err := c.Bind(&uaDto)
	// //if err != nil
	clientId := c.Get("uid").(uint64)
	ac.svc.upgradeAct(c.Request().Context(), clientId, dto.UpgradeActDto{})
	return nil
}
