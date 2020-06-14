package main

import "fmt"

func generate(n int, k int, start int, c []int, res *[][]int) {
	if len(c) == k {
		s := make([]int, len(c))
		copy(s, c)
		*res = append(*res, s)
		return
	}
	for i := start; i <= n; i++ {
		c = append(c, i)
		generate(n, k, i + 1, c, res)
		c = c[:len(c) - 1]
	}
}

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	if n <= 0 || n < k{
		return res
	}
	generate(n, k, 1, []int{}, &res)
	return res
}

func main() {
	fmt.Println(combine(4, 3))
}
