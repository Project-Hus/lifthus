package rec

import (
	"net/http"
	"routine/common/db"
	"routine/common/dto"

	"github.com/labstack/echo/v4"
)

// createWeeklyProgramRec godoc
// @Router       /rec/program/weekly [post]
// @Param createWeeklyProgramRecDto body dto.CreateWeeklyProgramRecDto true "createWeeklyProgramRec DTO"
// @Summary      gets CreateWeeklyProgramRecDto from body and returns created rec's ID
// @Tags         rec
// @Success      201 "returns created rec's ID"
// @Failure      400 "invalid body"
// @Failure      401 "unauthorized"
// @Failure 	 403 "forbidden"
// @Failure      500 "failed to create rec"
func (rc recApiController) createWeeklyProgramRec(c echo.Context) error {
	createProgramRecDto := new(dto.CreateWeeklyProgramRecDto)
	if err := c.Bind(createProgramRecDto); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	aid := createProgramRecDto.Author
	uid := c.Get("uid").(uint64)
	if aid != uid {
		return c.String(http.StatusForbidden, "you are not allowed to create rec for others")
	}

	rid, err := db.CreateWeeklyProgramRec(rc.dbClient, c.Request().Context(), *createProgramRecDto)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, rid)
}
