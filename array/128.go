package main

import (
	"fmt"
	"sort"
)


func longestConsecutive(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	sort.Ints(nums)
	maxLength := 1
	tmpLength := 1
	start := nums[0]
	for i := 1; i < length; i++ {
		if nums[i] == start + 1 {
			tmpLength++
		}else if nums[i] == start {
			continue
		}else {
			if tmpLength > maxLength {
				maxLength = tmpLength
			}
			tmpLength = 1
		}
		start = nums[i]
	}
	if tmpLength > maxLength {
		maxLength = tmpLength
	}
	return maxLength
}

func main() {
	a := []int{1, 2, 0, 1}
	fmt.Println(longestConsecutive(a))
}
