package main

import (
	"fmt"
)

func bubbleSort(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		flag := false
		for j := 1; j < length - i;j++ {
			if nums[j - 1] > nums[j] {
				nums[j - 1], nums[j] = nums[j], nums[j - 1]
				flag = true
			}
		}
		if flag == false {
			return
		}
	}
}

func main() {
	a := []int{12, 1, 9, 21, 3, 20, 0}
	bubbleSort(a)
	fmt.Println(a)
}
