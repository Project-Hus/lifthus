package helper

import "github.com/golang-jwt/jwt/v5"

// SignedHusTotalSignOutToken generates it with given hsid.
func SignedHusTotalSignOutToken(hsid string) (st string, err error) {
	token := NewJWT(jwt.MapClaims{
		"pps":  "hus_signout",
		"hsid": hsid,
		"type": "total",
	})
	return SignJWT(token)
}
