package relation

import (
	"lifthus-auth/db"
	"lifthus-auth/ent"
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
// @Success      200 "signed user now follows the given user"
// @Failure      400 "invalid uid"
// @Failure      404 "user not found"
// @Failure      500 "failed to get user following list"
func (rc relationApiController) FollowUser(c echo.Context) error {
	signedUser := c.Get("uid").(uint64)
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

	// try adding the follower
	_, err = u.Update().AddFollowerIDs(signedUser).Save(c.Request().Context())
	if err != nil {
		// maybe already following
		if ent.IsConstraintError(err) {
			// then try unfollowing
			_, err = u.Update().RemoveFollowerIDs(signedUser).Save(c.Request().Context())
			if err == nil {
				// unfollowed successfully
				goto SUCC
			}
		}
		return c.String(http.StatusInternalServerError, err.Error())
	}
SUCC:
	return c.String(http.StatusOK, "followed or unfollowed successfully")
}
