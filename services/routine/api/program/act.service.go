package program

import "github.com/labstack/echo/v4"

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
	return nil
}
