package program

import (
	"net/http"
	"routine/common/db"
	"routine/common/dto"

	"github.com/labstack/echo/v4"
)

// createWeeklyProgram godoc
// @Router       /routine/program/weekly [post]
// @Param createWeeklyProgramDto body dto.CreateWeeklyProgramDto true "createWeeklyProgram DTO"
// @Summary      gets CreateWeeklyProgramDto from body and returns created program
// @Tags         program
// @Success      201 "returns created program"
// @Failure      400 "invalid body"
// @Failure      401 "unauthorized"
// @Failure 	 403 "forbidden"
// @Failure      500 "failed to create program"
func (rc programApiController) createWeeklyProgram(c echo.Context) error {
	createProgramDto := new(dto.CreateWeeklyProgramDto)
	if err := c.Bind(createProgramDto); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	createdProgram, err := db.CreateWeeklyProgram(rc.dbClient, c.Request().Context(), createProgramDto)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, createdProgram)
}
