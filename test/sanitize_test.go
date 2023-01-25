package test

import (
	"packages-api/utils"
	"testing"
)

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
		{
			name:     "remove whitespaces from empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "remove whitespaces from leading whitespaces string",
			input:    "   leading whitespaces",
			expected: "leadingwhitespaces",
		},
		{
			name:     "remove whitespaces from trailing whitespaces string",
			input:    "trailing whitespaces   ",
			expected: "trailingwhitespaces",
		},
		{
			name:     "remove whitespaces from mixed string",
			input:    "   mixed   whitespaces   ",
			expected: "mixedwhitespaces",
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

func TestRemoveDots(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "remove dots",
			input:    "Hello..World.",
			expected: "HelloWorld",
		},
		{
			name:     "remove dots from single dot string",
			input:    ".",
			expected: "",
		},
		{
			name:     "remove dots from multiple dot string",
			input:    ".......",
			expected: "",
		},
		{
			name:     "remove dots from mixed string",
			input:    "Hello World..This is. a test.",
			expected: "HelloWorldThisisatest",
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

func TestRemoveNonAlphanumeric(t *testing.T) {
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
		{
			name:     "remove non-alphanumeric characters from string with only special characters",
			input:    "@#$%^&*",
			expected: "",
		},
		{
			name:     "remove non-alphanumeric characters from string with mixed characters",
			input:    "Hello World! This is a test.",
			expected: "HelloWorldThisisatest",
		},
		{
			name:     "remove non-alphanumeric characters from string with numbers",
			input:    "Hello World! This is a test. 1 2 3",
			expected: "HelloWorldThisisatest123",
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
		{
			name:     "convert to lowercase with mixed cases",
			input:    "HeLLo WorlD. This Is A Test",
			expected: "hello world. this is a test",
		},
		{
			name:     "convert to lowercase with special characters",
			input:    "HeLLo WorlD. This Is A Test #$%^&*",
			expected: "hello world. this is a test #$%^&*",
		},
		{
			name:     "convert to lowercase with numbers",
			input:    "HeLLo WorlD. This Is A Test #$%^&*123",
			expected: "hello world. this is a test #$%^&*123",
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
