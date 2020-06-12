package main

import "sort"

func threeSum(nums []int) [][]int {
	length := len(nums)
	var res [][]int
	if length <= 1 {
		return res
	}
	sort.Ints(nums)
	if nums[0] <= 0 && nums[length - 1] >= 0 {

	}
}


func main() {
	
}
