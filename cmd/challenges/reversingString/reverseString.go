package main

import "fmt"

func reverseString(str string) string {
	chars := []rune(str)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func main() {
	n := "Hello, world"

	fmt.Printf("reverseString(n): %v\n", reverseString(n))
}
