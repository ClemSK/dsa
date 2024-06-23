package main

import (
	"fmt"
	"testing"
)

func TestReverseStr(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"world", "dlrow"},
		{"12345", "54321"},
		{"!@#$%", "%$#@!"},
		{"racecar", "racecar"},
		{"こんにちは", "はちにんこ"},
	}

	for _, test := range tests {
		output := reverseString(test.input)
		fmt.Printf("Test input: %v, expected: %v, output: %v\n", test.input, test.expected, reverseString(test.input))
		if output != test.expected {
			t.Errorf("reverseString(%s) = %s; expected %s", test.input, output, test.expected)
		}
	}
}
