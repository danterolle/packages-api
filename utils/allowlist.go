package utils

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
