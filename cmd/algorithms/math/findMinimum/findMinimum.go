package main

import (
	"fmt"
	"math"
)

func findMinimum(nums []int) int {
	if len(nums) == 0 {
		return math.MaxInt64
	}

	min := math.MaxInt64
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func main() {
	n := []int{1, 4, 6, 8, 23}

	fmt.Printf("findMinimum(%d) = %d\n", n, findMinimum(n))
}
