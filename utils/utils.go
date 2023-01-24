package utils

import (
	"regexp"
	"strings"
)

// SanitizeInput Sanitize the user input for the variables branch and arch
// to ensure that they do not contain any "../", "..", or similar sequences that could be used to
// navigate outside the intended directory.
func SanitizeInput(input string) string {
	// Remove any ".." or "." from the input string
	input = strings.Replace(input, "..", "", -1)
	input = strings.Replace(input, ".", "", -1)

	// Remove any leading or trailing whitespaces
	input = strings.TrimSpace(input)

	// Remove any non-alphanumeric characters
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	input = reg.ReplaceAllString(input, "")

	// Convert input to lowercase
	input = strings.ToLower(input)

	return input
}
