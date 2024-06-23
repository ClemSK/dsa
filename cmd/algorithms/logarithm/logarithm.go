package main

import (
	"fmt"
	"math"
)

func logarithm(num_followers int, average_engagement_percentage float64) float64 {
	if num_followers <= 0 {
		return 0
	}
	return math.Ceil(average_engagement_percentage * math.Log2(float64(num_followers)))
}

func main() {
	n := 4000
	p := 0.3

	fmt.Printf("logarithm(%v, %v) = %v\n", n, p, logarithm(n, p))
}
