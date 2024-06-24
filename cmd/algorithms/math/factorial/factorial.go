package main

import (
	"fmt"
	"time"
)

func factorialLoop(numPosts int) int {
	if numPosts <= 0 {
		return 1
	}
	start1 := time.Now()
	posts := 1
	for i := 1; i <= numPosts; i++ {
		posts *= i
	}

	executionTime1 := time.Since(start1)
	fmt.Printf("posts: %v, executionTime1: %v\n", posts, executionTime1)
	return posts
}

func factorialRecursion(numPosts int) int {
	if numPosts <= 0 {
		return 1
	}
	return numPosts * factorialRecursion(numPosts-1)
}

func main() {
	n := 20

	start1 := time.Now()
	result1 := factorialLoop(n)
	executionTime1 := time.Since(start1)

	start2 := time.Now()
	result2 := factorialRecursion(n)
	executionTime2 := time.Since(start2)

	fmt.Printf("factorialLoop(n): %v, execution time: %v\n", result1, executionTime1)
	// factorialLoop(n): 2432902008176640000, execution time: 22.708µs

	fmt.Printf("factorialRecursion(n): %v, execution time: %v\n", result2, executionTime2)
	// factorialRecursion(n): 2432902008176640000, execution time: 42ns
}
