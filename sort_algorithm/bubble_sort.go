package main

import (
	"fmt"
)

func bubbleSort(nums []int) {
	length := len(nums)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length;j++ {
			if nums[j] < nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}
}

func main() {
	a := []int{12, 1, 9, 21, 3, 20, 0}
	bubbleSort(a)
	fmt.Println(a)
}
