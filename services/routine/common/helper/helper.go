package helper

import (
	"net/url"
	"strings"
)

// Slugify returns slugified string
func Slugify(str string) string {
	specialChars := []string{"@", "#", "$", "%", "&", "*", "?"}

	// replace all spaces of str with '-'
	slug := strings.ReplaceAll(str, " ", "-")

	for i := 0; i < len(specialChars); i++ {
		encodedChar := url.QueryEscape(specialChars[i])

		// split slug by specialChars[i] and join it with encodedChar
		slug = strings.Join(strings.Split(slug, specialChars[i]), encodedChar)
	}

	return slug
}
