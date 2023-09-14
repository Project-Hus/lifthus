package program

import (
	"net/http"
	"routine/ent/act"
	"routine/shared/db"
	"routine/shared/dto"
	"strconv"

	"github.com/labstack/echo/v4"
)

// queryAct godoc
// @Router       /act [get]
// @Param name query string false "act name"
// @Param skip query int false "skip"
// @Param id query int false "act id"
// @Summary      gets Act name from query-string and returns corresponding Acts
// @Tags         act
// @Success      200 "returns acts"
// @Failure      400 "invalid request"
// @Failure      500 "failed to query acts"
func (pc programApiController) queryAct(c echo.Context) error {
	// if ID is given, it has priority over name
	actId := c.QueryParam("id")
	if actId != "" {
		id, err := strconv.ParseUint(actId, 10, 64)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		act, err := pc.dbClient.Act.Query().Where(act.IDEQ(id)).Only(c.Request().Context())
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, act)
	}
	// from query-string get name and skip
	actName := c.QueryParam("name")
	skipStr := c.QueryParam("skip")
	// convert skip to int if it exists
	skip, err := strconv.Atoi(skipStr)
	if skipStr == "" {
		skip = 0
	} else {
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
	}
	acts, err := db.QueryActsByName(pc.dbClient, c.Request().Context(), actName, skip)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, acts)
}

// createAct godoc
// @Router       /act [post]
// @Param createActDto body dto.CreateActDto true "createAct DTO"
// @Summary      gets CreateActDto from body and returns created act's ID
// @Tags         act
// @Success      201 "returns created act's ID"
// @Failure      400 "invalid body"
// @Failure      401 "unauthorized"
// @Failure 	 403 "forbidden"
// @Failure      500 "failed to create program"
func (pc programApiController) createAct(c echo.Context) error {
	createActDto := new(dto.CreateActDto)
	if err := c.Bind(createActDto); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	aid := createActDto.Author
	uid := c.Get("uid").(uint64)
	if aid != uid {
		return c.String(http.StatusForbidden, "you are not allowed to create program for others")
	}

	aid, err := db.CreateAct(pc.dbClient, c.Request().Context(), createActDto)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, aid)
}
