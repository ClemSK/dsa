package main

import "fmt"

func sumAll(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	n := []int{1, 4, 6, 8, 23}

	fmt.Printf("sumAll(%d) = %d\n", n, sumAll(n))
}
