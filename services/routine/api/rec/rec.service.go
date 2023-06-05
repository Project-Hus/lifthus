package rec

import (
	"net/http"
	"routine/common/db"
	"routine/common/dto"
	"routine/ent"
	"routine/ent/dailyroutinerec"
	"time"

	"github.com/labstack/echo/v4"
)

// updateRoutineActRec godoc
// @Router       /rec/routineact [put]
// @param updateRoutineActRecDto body dto.UpdateRoutineActRec true "update routineact rec dto"
// @Summary      updates routineact rec
// @Tags         rec
// @Success      200 "returns updated routineact rec"
// @Failure      400 "invalid request"
// @Failure      401 "unauthorized"
// @Failure      403 "forbidden"
// @Failure      500 "failed to update routineact rec"
func (rc recApiController) updateRoutineActRec(c echo.Context) error {
	return nil
}

// queryRoutineActRecs godoc
// @Router       /rec/routineact [get]
// @Param date query string false "date like 2006-01-02"
// @Param startDate query string false "start date like 2006-01-02"
// @Param endDate query string false "end date like 2006-01-02s"
// @Summary      gets specific date or range of date and returns routineact recs
// @Tags         rec
// @Success      200 "returns routineact recs"
// @Failure      400 "invalid request"
// @Failure      401 "unauthorized"
// @Failure 	 403 "forbidden"
// @Failure      500 "failed to query routineact recs"
func (rc recApiController) queryRoutineActRecs(c echo.Context) error {
	uid := c.Get("uid").(uint64)

	dateQ := c.QueryParam("date")
	// if date is given, handle specific date query
	if dateQ != "" {
		date, err := time.Parse("2006-01-02", dateQ)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		rars, err := rc.dbClient.DailyRoutineRec.Query().
			Where(dailyroutinerec.DateEQ(date), dailyroutinerec.AuthorEQ(uid)).
			WithRoutineActRecs().
			All(c.Request().Context())
		if ent.IsNotFound(err) {
			rars = []*ent.DailyRoutineRec{}
			return c.JSON(http.StatusOK, rars)
		} else if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, rars)
	}
	// date is not given, handle range of date query
	startDateQ := c.QueryParam("startDate")
	endDateQ := c.QueryParam("endDate")
	if startDateQ == "" || endDateQ == "" {
		return c.String(http.StatusBadRequest, "invalid request")
	}
	startDate, errS := time.Parse("2006-01-02", startDateQ)
	endDate, errE := time.Parse("2006-01-02", endDateQ)
	if errS != nil || errE != nil {
		return c.String(http.StatusBadRequest, "invalid request")
	}
	rars, err := rc.dbClient.DailyRoutineRec.Query().
		Where(dailyroutinerec.DateGTE(startDate), dailyroutinerec.DateLTE(endDate), dailyroutinerec.AuthorEQ(uid)).
		WithRoutineActRecs().
		All(c.Request().Context())
	if ent.IsNotFound(err) {
		rars = []*ent.DailyRoutineRec{}
		return c.JSON(http.StatusOK, rars)
	} else if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rars)
}

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
