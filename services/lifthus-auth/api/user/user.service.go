package user

import (
	"fmt"
	"lifthus-auth/common/dto"
	"lifthus-auth/db"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetUserInfo godoc
// @Router       /auth/user/{uid} [get]
// @Param uid path string true "user id"
// @Summary      gets uid from path param and returns user info
// @Description  if the signed user is the same as the requested user, returns all info while hiding sensitive info if different.
// @Tags         user
// @Success      200 "returns user info as json"
// @Failure      400 "invalid uid"
// @Failure      404 "user not found"
// @Failure      500 "failed to create new session"
func (uc userApiController) GetUserInfo(c echo.Context) error {
	signedUser, ok := c.Get("uid").(uint64)

	uid, err := strconv.ParseUint(c.Param("uid"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user, err := db.QueryUserByUID(c.Request().Context(), uc.dbClient, uint64(uid))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if user == nil {
		return c.String(http.StatusNotFound, fmt.Sprintf("user %d not found", uid))
	}

	var udto *dto.QueryUserDto

	if ok && signedUser == uid {
		udto = &dto.QueryUserDto{
			ID:              user.ID,
			Registered:      user.Registered,
			RegisteredAt:    user.RegisteredAt,
			Username:        user.Username,
			Email:           &user.Email,
			EmailVerified:   &user.EmailVerified,
			Name:            &user.Name,
			GivenName:       &user.GivenName,
			FamilyName:      &user.FamilyName,
			Birthdate:       user.Birthdate,
			ProfileImageURL: *user.ProfileImageURL,
			CreatedAt:       user.CreatedAt,
			UpdatedAt:       user.UpdatedAt,
		}
	} else {
		udto = &dto.QueryUserDto{
			ID:              user.ID,
			Registered:      user.Registered,
			RegisteredAt:    user.RegisteredAt,
			Username:        user.Username,
			Email:           nil,
			EmailVerified:   nil,
			Name:            nil,
			GivenName:       nil,
			FamilyName:      nil,
			Birthdate:       user.Birthdate,
			ProfileImageURL: *user.ProfileImageURL,
			CreatedAt:       user.CreatedAt,
			UpdatedAt:       user.UpdatedAt,
		}
	}

	return c.JSON(http.StatusOK, udto)
}
