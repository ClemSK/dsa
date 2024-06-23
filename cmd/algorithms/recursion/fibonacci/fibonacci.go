package main

import "fmt"

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &n)
	fmt.Printf("Fibonacci(%d) = %d\n", n, fib(n))
}
