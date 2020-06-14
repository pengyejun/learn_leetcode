package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b := 1, 1
	for n > 0 {
		if n >= 2 {
			a, b = a + b, a
		}
		n -= 1
	}
	return a
}

func main() {
	fmt.Println(climbStairs(4))
}
