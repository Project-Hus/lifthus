package helper

import (
	"fmt"
	"lifthus-auth/common/lifthus"
	"lifthus-auth/ent"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

// SignedLST generates it with given session.
// Don't forget to include User edges.
func SignedLST(ls *ent.Session) (st string, err error) {
	token := NewJWT(jwt.MapClaims{
		"pps": "lifthus_session",
		"iss": "https://auth.lifthus.com",
		"sid": ls.ID.String(),
		"tid": ls.Tid.String(),
		"uid": UIDToString(ls.Edges.User),
		"exp": lifthus.GetLstExp(),
	})
	return SignJWT(token)
}

// SignedHusConnectionToken generates it with given lifthus sessoin ID.
func SignedHusConnectionToken(lsid string) (hct string, err error) {
	ct := NewJWT(jwt.MapClaims{
		"pps":     "hus_connection",
		"iss":     "https://auth.lifthus.com",
		"service": "lifthus",
		"sid":     lsid,
		"exp":     time.Now().Add(time.Second * 10).Unix(),
	})
	return SignJWT(ct)
}

// SignedHusTotalSignOutToken generates it with given hsid.
func SignedHusTotalSignOutToken(hsid string) (st string, err error) {
	token := NewJWT(jwt.MapClaims{
		"pps":  "hus_signout",
		"hsid": hsid,
		"type": "total",
	})
	return SignJWT(token)
}
