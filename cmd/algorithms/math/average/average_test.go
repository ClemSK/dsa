package main

import (
	"fmt"
	"testing"
)

func TestAverage(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{0, 0, 0, 0, 0}, 0},
		{[]int{-1, -2, -3, -4, -5}, -3},
		{[]int{1, 4, 6, 8, 23}, 8},
		{[]int{10, 200, 3000, 5000, 4}, 1642},
		{[]int{12, 12, 12}, 12},
		{[]int{7, 4, 3, 100, 2343243, 343434, 1, 2, 32}, 298536},
	}

	for _, test := range tests {
		output := average(test.input)
		fmt.Printf("Test input: %v, expected: %v, output: %v\n", test.input, test.expected, average(test.input))
		if output != test.expected {
			t.Errorf("average(%v) = %d; expected %d", test.input, output, test.expected)
		}
	}
}
