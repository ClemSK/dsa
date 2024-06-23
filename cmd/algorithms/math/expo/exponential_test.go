package main

import (
	"fmt"
	"math"
	"testing"
)

func TestExponential(t *testing.T) {
	tests := []struct {
		input    []int
		expected float64
	}{
		{[]int{2, 3, 2, 19}, 324.25103738467675},
		{[]int{}, 0},
		{[]int{5, 5, 5, 5}, 182.05642030260802},
		{[]int{1}, 1},
		// {[]int{-1, -2, -3, -4}, 11.097034831066584},
	}

	for _, test := range tests {
		output := exponential(test.input)
		fmt.Printf("Test input: %v, expected: %v, output: %v\n", test.input, test.expected, exponential(test.input))
		if !floatEquals(output, test.expected) {
			t.Errorf("exponential(%v) = %.6f; expected %.6f", test.input, output, test.expected)
		}
	}
}

func floatEquals(a, b float64) bool {
	const epsilon = 1e-6
	return math.Abs(a-b) < epsilon
}
