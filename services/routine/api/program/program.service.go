package program

import (
	"net/http"
	"routine/common/db"
	"routine/common/dto"

	"github.com/labstack/echo/v4"
)

// createWeeklyProgram godoc
// @Router       /program/weekly [post]
// @Param createWeeklyProgramDto body dto.CreateWeeklyProgramDto true "createWeeklyProgram DTO"
// @Summary      gets CreateWeeklyProgramDto from body and returns created program's ID
// @Tags         program
// @Success      201 "returns created program's ID"
// @Failure      400 "invalid body"
// @Failure      401 "unauthorized"
// @Failure 	 403 "forbidden"
// @Failure      500 "failed to create program"
func (pc programApiController) createWeeklyProgram(c echo.Context) error {
	createProgramDto := new(dto.CreateWeeklyProgramDto)
	if err := c.Bind(createProgramDto); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	aid := createProgramDto.Author
	uid := c.Get("uid").(uint64)
	if aid != uid {
		return c.String(http.StatusForbidden, "you are not allowed to create program for others")
	}

	pid, err := db.CreateWeeklyProgram(pc.dbClient, c.Request().Context(), createProgramDto)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, pid)
}
