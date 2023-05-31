package helper

import "lifthus-auth/common/helper"

func ValidateLST(lst string) (uid *uint64, err error) {
	lstClaims, exp, err := helper.ParseJWTwithHMAC(lst)
	return uid, err
}

// "purpose": "lifthus_session",
// 		"sid":     sid,
// 		"uid":     strconv.FormatUint(*ls.UID, 10),
// 		"exp":     time.Now().Add(time.Minute * 5).Unix(),
