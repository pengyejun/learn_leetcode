package main

import "fmt"

func mergeSort(nums []int) []int{
	if len(nums) <= 1 {
		return nums
	}
	n := len(nums) / 2
	left := mergeSort(nums[:n])
	right := mergeSort(nums[n:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	l, r := 0, 0
	result := make([]int, 0)
	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			result = append(result, right[r])
			r++
		}else {
			result = append(result, left[l])
			l++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}

func main() {
	a := []int{12, 1, 9, 21, 3, 20, 0}
	fmt.Println(mergeSort(a))
	fmt.Println(a)
}
