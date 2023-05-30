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
// @Router       /user/{uid} [get]
// @Param uid path string true "user id"
// @Summary      gets uid from path param and returns user info
// @Description  if the signed user is the same as the requested user, returns all info while hiding sensitive info if different.
// @Tags         user
// @Success      200 "returns user info as json"
// @Failure      400 "invalid uid"
// @Failure      404 "user not found"
// @Failure      500 "failed to get user info"
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

			Usercode: user.Usercode,
			Company:  user.Company,
			Location: user.Location,
			Contact:  user.Contact,
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

			Usercode: user.Usercode,
			Company:  user.Company,
			Location: user.Location,
			Contact:  user.Contact,
		}
	}

	return c.JSON(http.StatusOK, udto)
}

// GetUserInfoByUsername godoc
// @Router       /username/{username} [get]
// @Param username path string true "user id"
// @Summary      gets username from path param and returns user info
// @Description  if the signed user is the same as the requested user, returns all info while hiding sensitive info if different.
// @Tags         user
// @Success      200 "returns user info as json"
// @Failure      404 "user not found"
// @Failure      500 "failed to get user info"
func (uc userApiController) GetUserInfoByUsername(c echo.Context) error {
	signedUser, ok := c.Get("uid").(uint64)

	username := c.Param("username")

	user, err := db.QueryUserByUsername(c.Request().Context(), uc.dbClient, username)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if user == nil {
		return c.String(http.StatusNotFound, fmt.Sprintf("user %s not found", username))
	}

	var udto *dto.QueryUserDto

	if ok && signedUser == user.ID {
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

			Usercode: user.Usercode,
			Company:  user.Company,
			Location: user.Location,
			Contact:  user.Contact,
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

			Usercode: user.Usercode,
			Company:  user.Company,
			Location: user.Location,
			Contact:  user.Contact,
		}
	}

	return c.JSON(http.StatusOK, udto)
}

// SetUserInfo godoc
// @Router       /user [put]
// @Param userinfo body dto.UpdateUserInfoDto true "user info"
// @Summary      gets uid from path param and updates user info
// @Description  it gets uid from path param and updates user info
// @Tags         user
// @Success      200 "returns user info as json"
// @Failure      400 "invalid uid"
// @Failure      404 "user not found"
// @Failure      500 "failed to set user info"
func (uc userApiController) SetUserInfo(c echo.Context) error {
	uid := c.Get("uid").(uint64)
	userInfo := new(dto.UpdateUserInfoDto)
	if err := c.Bind(userInfo); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	if userInfo.Uid != uid {
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}
	user, err := db.UpdateUserInfo(c.Request().Context(), uc.dbClient, *userInfo)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	qu := &dto.QueryUserDto{
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

		Usercode: user.Usercode,
		Company:  user.Company,
		Location: user.Location,
		Contact:  user.Contact,
	}

	return c.JSON(http.StatusOK, qu)
}
