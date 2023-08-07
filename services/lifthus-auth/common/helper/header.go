package helper

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

// CheckAuthHeader checks if given authorization header value has valid format.
func CheckAuthHeader(ah string) (token string, err error) {
	if ah != "" && !strings.HasPrefix(ah, "Bearer ") {
		return "", fmt.Errorf("invalid authorization header")
	}
	return strings.TrimPrefix(ah, "Bearer "), nil
}

// GenAuthHeader generates authorization header value with given token.
func GenAuthHeader(token string) string {
	return fmt.Sprintf("Bearer %s", token)
}

// SetAuthHeader sets Authorization header to given context with given token.
func SetAuthHeader(c echo.Context, token string) echo.Context {
	c.Response().Header().Set("Authorization", GenAuthHeader(token))
	return c
}