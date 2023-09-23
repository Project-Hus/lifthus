package domain

import (
	"crypto/rand"
	"encoding/hex"
)

// RandomHexCode generates random hex code that is expected to be unique in some part of the system.
func RandomHexCode() (string, error) {
	return RandomHex(CODE_LENGTH)
}

// RandomHex generates random hex string with given length.
func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
