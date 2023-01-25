package utils

import (
	"regexp"
	"strings"
	"unicode"
)

func RemoveDots(input string) string {
	// Remove any ".." or "." from the input string
	input = strings.Replace(input, "..", "", -1)
	input = strings.Replace(input, ".", "", -1)

	return input
}

func RemoveWhitespaces(input string) string {
	reg, err := regexp.Compile("\\s+")
	if err != nil {
		return input
	}
	return reg.ReplaceAllString(input, "")
}

func RemoveNonAlphanumeric(input string) string {
	return strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return r
	}, input)
}

func ConvertToLowercase(input string) string {
	input = strings.ToLower(input)

	return input
}

// SanitizeInput Sanitize the user input for the variables branch and arch
// to ensure that they do not contain any "../", "..", or similar sequences that could be used to
// navigate outside the intended directory.
func SanitizeInput(input string) string {
	// Remove any ".." or "." from the input string
	RemoveDots(input)

	// Remove any leading or trailing whitespaces
	RemoveWhitespaces(input)

	// Remove any non-alphanumeric characters
	RemoveNonAlphanumeric(input)

	// Convert input to lowercase
	ConvertToLowercase(input)

	return input
}

// CheckAllowlist it is used for checking if a given input string is present in both Branch and Arch constants
func CheckAllowlist(input string, allowlist []string) bool {
	// Iterate over the allowlist values (check constants.go to see the values allowed)
	for _, value := range allowlist {
		// Check if the input value is equal to the current allowlist value
		if input == value {
			return true
		}
	}
	// If the input value is not found in the allowlist, return false
	return false
}
