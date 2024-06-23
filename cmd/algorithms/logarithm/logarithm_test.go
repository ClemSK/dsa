package main

import (
	"fmt"
	"math"
	"testing"
)

func TestLogarithm(t *testing.T) {
	tests := []struct {
		num_followers                 int
		average_engagement_percentage float64
		expected                      float64
	}{
		{1000, 10.0, 100},    // Expected output rounded up
		{100, 5.0, 34},       // Expected output rounded up
		{1, 10.0, 0},         // Log2(1) is 0, so result should be 0
		{10000, 7.5, 100},    // Expected output rounded up
		{500, 2.5, 23},       // Expected output rounded up
		{1000000, 20.0, 399}, // Expected output rounded up
		{0, 10.0, 0},         // Invalid input, should return 0
		{-100, 5.0, 0},       // Invalid input, should return 0
	}

	for _, test := range tests {
		output := logarithm(test.num_followers, test.average_engagement_percentage)
		fmt.Printf("num_followers: %v, average_engagement_percentage: %v, expected: %v, output: %v\n", test.num_followers, test.average_engagement_percentage, test.expected, logarithm(test.num_followers, test.average_engagement_percentage))
		if !floatEquals(output, test.expected) {
			t.Errorf("logarithm(%d, %.1f) = %.6f; expected %.6f", test.num_followers, test.average_engagement_percentage, output, test.expected)
		}
	}
}

func floatEquals(a, b float64) bool {
	const epsilon = 1e-6
	return math.Abs(a-b) < epsilon
}
