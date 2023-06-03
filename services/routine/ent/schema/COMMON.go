package schema

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// enum
var programType = []string{"weekly", "daily"}
var actType = []string{"rep", "lap", "simple"}
var recStatus = []string{"history", "waiting", "proceeding", "completed", "failed", "canceled"}

// func
func randomHex() string {
	bytes := make([]byte, 4)
	if _, err := rand.Read(bytes); err != nil {
		fmt.Println(err)
		return ""
	}
	return hex.EncodeToString(bytes)
}
