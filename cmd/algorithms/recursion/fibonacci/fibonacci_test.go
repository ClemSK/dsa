package main

import (
	"fmt"
	"testing"
)

// TestFib tests the Fib function.
func TestFib(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{10, 55},
		{20, 6765},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Fib(%d)", tc.n), func(t *testing.T) {
			result := fib(tc.n)
			if result != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, result)
			}
		})
	}
}
