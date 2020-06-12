package main

func nextGreaterElements(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	stack := make([]int, 0)
	for i := 0; i < length; i++ {
		res[i] = -1
		for len(stack) > 0 && nums[stack[len(stack) - 1]] < nums[i] {
			res[stack[len(stack) - 1]] =nums[i]
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, i)
	}

	for i := 0; i < length; i++ {
		for len(stack) > 0 && nums[stack[len(stack) - 1]] < nums[i] {
			res[stack[len(stack) - 1]] = nums[i]
			stack = stack[:len(stack) - 1]
		}
	}
	return res
}

func main() {
	nextGreaterElements([]int{1, 2, 1})
}
