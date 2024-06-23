package main

import (
	"fmt"
	"math"
)

func exponential(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}
	var sum int
	for _, num := range nums {
		sum += num
	}
	averageAudienceFollowers := float64(sum) / float64((len(nums)))

	est_spread := averageAudienceFollowers * math.Pow(float64(sum), 1.2)

	return est_spread
}

func main() {
	n := []int{2, 3, 2, 19}
	// n1 := []int{}
	// n2 := []int{5, 5, 5, 5}
	// n3 := []int{1}

	fmt.Printf("expo(%v): %v\n", n, exponential(n))
	// fmt.Printf("expo(%v): %v\n", n1, exponential(n1))
	// fmt.Printf("expo(%v): %v\n", n2, exponential(n2))
	// fmt.Printf("expo(%v): %v\n", n3, exponential(n3))
}
