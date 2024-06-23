package main

import (
	"fmt"
	"testing"
)

func TestSumAll(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{150, 200, 50, 300, 100}, 800},
		{[]int{}, 0}, // Empty slice should return 0
		{[]int{5, 4, 3, 2, 1}, 15},
		{[]int{-1, -5, -3, -2, -4}, -15},
		{[]int{1, 4, 6, 8, 23}, 42},           // Sum of {1, 4, 6, 8, 23} is 42
		{[]int{10, 200, 3000, 5000, 4}, 8214}, // Sum of {10, 200, 3000, 5000, 4} is 8214
		{[]int{12, 12, 12}, 36},               // Sum of {12, 12, 12} is 36
		{[]int{7, 4, 3, 100, 2343243, 343434, 1, 2, 32}, 2686826},
	}

	for _, test := range tests {
		output := sumAll(test.input)
		fmt.Printf("Test input: %v, expected: %v, output: %v\n", test.input, test.expected, sumAll(test.input))
		if output != test.expected {
			t.Errorf("sumAll(%v) = %d; expected %d", test.input, output, test.expected)
		}
	}
}
