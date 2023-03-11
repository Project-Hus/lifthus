package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// NewSessionHandler godoc
// @Router       /session/new [post]
// @Summary      gets new connection and assign a lifthus session token.
// @Description  when user connects to lifthus newly, the client requests new session token from root componenet.
// @Tags         auth
// @Success      201 "returns session id with session token in cookie"
// @Failure      501 "failed to create new session"
func (ac authApiController) NewSessionHandler(c echo.Context) error {

	return c.String(http.StatusCreated, "")
}
