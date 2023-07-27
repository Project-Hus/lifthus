package helper

import (
	"crypto/rand"
	"encoding/hex"
	"lifthus-auth/ent"
	"strconv"
)

// RandomHex generates random hex string with given length.
func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// UIDToString converts UID to string if exists. else return empty string.
func UIDToString(u *ent.User) string {
	if u == nil {
		return ""
	}
	return strconv.FormatUint(u.ID, 10)
}
