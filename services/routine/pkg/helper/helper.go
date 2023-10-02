package helper

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"
)

// Slugify returns slugified string
func Slugify(str string) string {
	specialChars := []string{"@", "#", "$", "&", "*", "?"}

	// replace all spaces of str with '-'
	slug := strings.ReplaceAll(str, " ", "-")

	for i := 0; i < len(specialChars); i++ {
		encodedChar := url.QueryEscape(specialChars[i])

		// split slug by specialChars[i] and join it with encodedChar
		slug = strings.Join(strings.Split(slug, specialChars[i]), encodedChar)
	}

	return slug
}

// RandomHex returns random hex string which length is l bytes
func RandomHex(l int) string {
	bytes := make([]byte, l)
	if _, err := rand.Read(bytes); err != nil {
		fmt.Println(err)
		return ""
	}
	return hex.EncodeToString(bytes)
}

func TrimSlash(s string) string {
	s = strings.TrimSuffix(s, "/")
	s = strings.TrimPrefix(s, "/")
	return s
}
