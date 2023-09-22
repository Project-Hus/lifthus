package domain

import (
	"crypto/rand"
	"encoding/hex"
)

// RandomHex generates random hex string with given length.
func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
