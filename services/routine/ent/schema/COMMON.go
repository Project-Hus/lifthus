package schema

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// enum
var ProgramType = []string{"weekly", "daily"}
var ActType = []string{"rep", "lap", "simple"}
var RecStatus = []string{"history", "waiting", "proceeding", "completed", "failed", "canceled"}

// func
func randomHex() string {
	bytes := make([]byte, 4)
	if _, err := rand.Read(bytes); err != nil {
		fmt.Println(err)
		return ""
	}
	return hex.EncodeToString(bytes)
}
