package helper

import (
	"fmt"
	"lifthus-auth/common/helper"
	"strconv"
)

func ValidateLST(lst string) (uid *uint64, err error) {
	lstClaims, exp, err := helper.ParseJWTWithHMAC(lst)
	if err != nil {
		return nil, err
	}
	if exp {
		return nil, nil
	}
	purpose, ok := lstClaims["pps"].(string)
	if !ok || purpose != "lifthus_session" {
		return nil, fmt.Errorf("invalid token")
	}
	uidString, ok := lstClaims["uid"].(string)
	if !ok {
		return nil, nil
	}
	uidInt, err := strconv.ParseUint(uidString, 10, 64)
	uid = &uidInt
	return uid, err
}
