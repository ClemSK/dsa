package main

import (
	"fmt"
	"math"
	"testing"
)

func TestFindMinimum(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{150, 200, 50, 300, 100}, 50},
		{[]int{}, math.MaxInt64},
		{[]int{5, 4, 3, 2, 1}, 1},
		{[]int{-1, -5, -3, -2, -4}, -5},
		{[]int{1, 4, 6, 8, 23}, 1},
		{[]int{10, 200, 3000, 5000, 4}, 4},
		{[]int{12, 12, 12}, 12},
		{[]int{7, 4, 3, 100, 2343243, 343434, 1, 2, 32}, 1},
	}

	for _, test := range tests {
		output := findMinimum(test.input)
		fmt.Printf("Test input: %v, expected: %v, output: %v\n", test.input, test.expected, findMinimum(test.input))
		if output != test.expected {
			t.Errorf("findMinimum(%v) = %d; expected %d", test.input, output, test.expected)
		}
	}
}
