package main

import "fmt"

func average(nums []int) int {
	var av int
	for _, num := range nums {
		av += num
	}

	return av / len(nums)
}

func main() {
	n := []int{1, 4, 6, 8, 23}

	fmt.Printf("average(%d) = %d\n", n, average(n))
}
