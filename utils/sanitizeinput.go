package utils

// SanitizeInput to sanitize the user input for the variables branch and arch
// to ensure that they do not contain any "../", "..", or similar sequences that could be used to
// navigate outside the intended directory.
// So, this function should avoid vulnerabilities like: /packages?branch=main&arch=amd64&package=../../../../../../root/.ssh/id_rsa
func SanitizeInput(input string) string {
	RemoveWhitespaces(input)
	RemoveDots(input)
	RemoveNonAlphanumeric(input)
	ConvertToLowercase(input)

	return input
}
