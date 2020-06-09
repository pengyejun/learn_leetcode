package main

import "fmt"

func translateNum(num int) int {
	a, b := 1, 1
	n := num % 10
	for num > 0 {
		num = num / 10
		m := num % 10
		tmp := 10 * m + n
		if 10 <= tmp && tmp <= 25 {
			a, b = a + b, a
		}else {
			b = a
		}
		n = m
	}
	return a
}

func main() {
	fmt.Println(translateNum(12258))
}
