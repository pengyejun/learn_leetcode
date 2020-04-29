package main

import "fmt"

func numTrees(n int) int {
	if n == 1 || n == 2{
		return n
	}
	m := map[int]int{0: 1, 1: 1, 2: 2}

	for i := 3; i <= n; i++ {
		tmp := 0
		for j := 1; j <= i; j++ {
			tmp += m[j - 1] * m[i - j]
		}
		m[i] = tmp
	}
	return m[n]
}


func main() {
	fmt.Println(numTrees(4))
}
