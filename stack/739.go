package main

func dailyTemperatures(T []int) []int {
	length := len(T)
	res := make([]int, length)
	stack := make([]int, 0)

	for i := length - 1; i >= 0; i-- {
		for len(stack) > 0 && T[i] >= T[stack[len(stack) - 1]] {
			stack = stack[:len(stack) - 1]
		}
		if len(stack) > 0 && T[i] < T[stack[len(stack) - 1]] {
			res[i] = stack[len(stack) - 1] - i
		}
		stack = append(stack, i)
	}
	return res
}

func main() {
	
}
