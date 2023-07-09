package relation

import (
	"lifthus-auth/common/db"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetUserFollowing godoc
// @Router       /relation/following/{uid} [get]
// @Param uid path string true "user id"
// @Summary      gets uid from path param and returns user's following list
// @Tags         relation
// @Success      200 "returns following list as list of number"
// @Failure      400 "invalid uid"
// @Failure      404 "user not found"
// @Failure      500 "failed to get user following list"
func (rc relationApiController) GetUserFollowing(c echo.Context) error {
	uid, err := strconv.ParseUint(c.Param("uid"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	u, err := db.QueryUserByUID(c.Request().Context(), rc.dbClient, uid)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	} else if u == nil {
		return c.String(http.StatusNotFound, "user not found")
	}

	following, err := u.QueryFollowing().IDs(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, following)
}

// GetUserFollowers godoc
// @Router       /relation/followers/{uid} [get]
// @Param uid path string true "user id"
// @Summary      gets uid from path param and returns user's follower list
// @Tags         relation
// @Success      200 "returns follower list as list of number"
// @Failure      400 "invalid uid"
// @Failure      404 "user not found"
// @Failure      500 "failed to get user follower list"
func (rc relationApiController) GetUserFollowers(c echo.Context) error {
	uid, err := strconv.ParseUint(c.Param("uid"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	u, err := db.QueryUserByUID(c.Request().Context(), rc.dbClient, uid)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	} else if u == nil {
		return c.String(http.StatusNotFound, "user not found")
	}

	followers, err := u.QueryFollowers().IDs(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, followers)
}

// FollowUser godoc
// @Router       /relation/follow/{uid} [post]
// @Param uid path string true "user id"
// @Summary      gets uid from path param and makes signed user follow the given user
// @Tags         relation
// @Success      200 "new following list"
// @Failure      400 "invalid uid"
// @Failure      404 "user not found"
// @Failure      500 "failed to get user following list"
func (rc relationApiController) FollowUser(c echo.Context) error {
	signedUid := c.Get("uid").(uint64)
	uid, err := strconv.ParseUint(c.Param("uid"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	signedUser, err := db.QueryUserByUID(c.Request().Context(), rc.dbClient, signedUid)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	} else if signedUser == nil {
		return c.String(http.StatusNotFound, "user not found")
	}

	// try adding the follower
	signedUser, err = signedUser.Update().AddFollowingIDs(uid).Save(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	following, err := signedUser.QueryFollowing().IDs(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, following)
}

// UnfollowUser godoc
// @Router       /relation/unfollow/{uid} [delete]
// @Param uid path string true "user id"
// @Summary      gets uid from path param and makes signed user unfollow the given user
// @Tags         relation
// @Success      200 "new following list"
// @Failure      400 "invalid uid"
// @Failure      404 "user not found"
// @Failure      500 "failed to get user following list"
func (rc relationApiController) UnfollowUser(c echo.Context) error {
	signedUid := c.Get("uid").(uint64)
	uid, err := strconv.ParseUint(c.Param("uid"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	signedUser, err := db.QueryUserByUID(c.Request().Context(), rc.dbClient, signedUid)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	} else if signedUser == nil {
		return c.String(http.StatusNotFound, "user not found")
	}

	// try adding the follower
	signedUser, err = signedUser.Update().RemoveFollowingIDs(uid).Save(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	following, err := signedUser.QueryFollowing().IDs(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, following)
}
