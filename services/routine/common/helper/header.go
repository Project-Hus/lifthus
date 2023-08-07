package helper

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

// GetHeaderLST gets the user's lifthus session token from the Authorization header.
func GetHeaderLST(c echo.Context) (string, error) {
	return CheckAuthHeader(c.Request().Header.Get("Authorization"))
}

// CheckAuthHeader checks if given authorization header value has valid format.
func CheckAuthHeader(ah string) (token string, err error) {
	if ah != "" && !strings.HasPrefix(ah, "Bearer ") {
		return "", fmt.Errorf("invalid authorization header")
	} else if ah == "" {
		return "", fmt.Errorf("no authorization header")
	}
	return strings.TrimPrefix(ah, "Bearer "), nil
}
