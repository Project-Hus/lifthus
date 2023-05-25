package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// RegisterUser godoc
// @Router       /user [post]
// @Param userinfo body RegisterInfoDto true "user register info"
// @Summary      gets user register info and registers user
// @Description  it gets register info and registers user to lifthus
// @Tags         user
// @Success      200 "returns register info as json"
// @Failure      400 "invalid body"
// @Failure      401 "unauthorized"
func (uc userApiController) RegisterUser(c echo.Context) error {
	uid := c.Get("uid").(uint64)
	// get RegisterInfoDto from body
	registerInfo := new(RegisterInfoDto)
	if err := c.Bind(registerInfo); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	if registerInfo.Uid != uid {
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}
	/* TODO: register info to Rec Service */
	return c.JSON(http.StatusOK, registerInfo)
}

type RegisterInfoDto struct {
	Uid          uint64  `json:"uid,omitempty"`
	TrainingType string  `json:"trainingType,omitempty"`
	BodyWeight   float64 `json:"bodyWeight,omitempty"`
	Height       float64 `json:"height,omitempty"`
	Squat        float64 `json:"squat,omitempty"`
	Benchpress   float64 `json:"benchpress,omitempty"`
	Deadlift     float64 `json:"deadlift,omitempty"`
}
