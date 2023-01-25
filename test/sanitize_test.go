package test

import (
	"packages-api/utils"
	"testing"
)

func TestRemoveDots(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "remove dots",
			input:    "Hello..World",
			expected: "HelloWorld",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := utils.RemoveDots(test.input)
			if actual != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, actual)
			}
		})
	}
}

func TestRemoveWhitespaces(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "remove whitespace",
			input:    "   Hello World   ",
			expected: "HelloWorld",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := utils.RemoveWhitespaces(test.input)
			if actual != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, actual)
			}
		})
	}
}

func TestRemoveChars(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "remove non-alphanumeric characters",
			input:    "He@llo World!",
			expected: "HelloWorld",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := utils.RemoveNonAlphanumeric(test.input)
			if actual != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, actual)
			}
		})
	}
}

func TestConvertToLowercase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "convert to lowercase",
			input:    "HeLLo WorlD",
			expected: "hello world",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := utils.ConvertToLowercase(test.input)
			if actual != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, actual)
			}
		})
	}
}
