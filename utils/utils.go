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
	reg := regexp.MustCompile(`^[\p{Ll}\p{Lu}\p{Nd}_-]+$`)
	input = reg.ReplaceAllString(input, "")

	// Convert input to lowercase
	input = strings.ToLower(input)

	return input
}

// CheckWhitelist it is used for checking if a given input string is present in both Branch and Arch constants
func CheckAllowlist(input string, whitelist []string) bool {
	// Iterate over the whitelist values (check constants.go to see the values allowed)
	for _, value := range whitelist {
		// Check if the input value is equal to the current whitelist value
		if input == value {
			return true
		}
	}
	// If the input value is not found in the whitelist, return false
	return false
}
