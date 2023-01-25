package utils

import (
	"regexp"
	"strings"
	"unicode"
)

// RemoveWhitespaces to remove any leading or trailing whitespaces
func RemoveWhitespaces(input string) string {
	reg, err := regexp.Compile("\\s+")
	if err != nil {
		return input
	}
	return reg.ReplaceAllString(input, "")
}

// RemoveDots to remove any ".." or "." from the input string
func RemoveDots(input string) string {
	input = RemoveWhitespaces(input)
	input = strings.Replace(input, "..", "", -1)
	input = strings.Replace(input, ".", "", -1)

	return input
}

// RemoveNonAlphanumeric to remove any non-alphanumeric characters
func RemoveNonAlphanumeric(input string) string {
	return strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return r
	}, input)
}

// ConvertToLowercase to convert input to lowercase
func ConvertToLowercase(input string) string {
	input = strings.ToLower(input)

	return input
}
