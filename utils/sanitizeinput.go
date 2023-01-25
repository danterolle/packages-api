package utils

// SanitizeInput Sanitize the user input for the variables branch and arch
// to ensure that they do not contain any "../", "..", or similar sequences that could be used to
// navigate outside the intended directory.
func SanitizeInput(input string) string {
	RemoveWhitespaces(input)
	RemoveDots(input)
	RemoveNonAlphanumeric(input)
	ConvertToLowercase(input)

	return input
}
