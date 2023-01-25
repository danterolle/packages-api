package test

import (
	"packages-api/utils"
	"testing"
)

func TestCheckAllowlist(t *testing.T) {
	tests := []struct {
		input     string
		allowlist []string
		expected  bool
	}{
		{"contrib", []string{"main", "contrib", "non-free"}, true},
		{"stable", []string{"main", "contrib", "non-free"}, false},
		{"main", []string{"main", "contrib", "non-free"}, true},
	}

	for _, test := range tests {
		actual := utils.CheckAllowlist(test.input, test.allowlist)
		if actual != test.expected {
			t.Errorf("Test failed: %s inputted, %v expected, received: %v", test.input, test.expected, actual)
		}
	}
}
