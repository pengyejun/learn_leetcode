package main

import "fmt"

func insertSort(nums []int) {
	length := len(nums)
	for i := 1; i < length; i++ {
		j := i - 1
		tmp := nums[i]
		for ;j >= 0 && nums[j] > tmp; j-- {
			nums[j + 1] = nums[j]
		}
		nums[j + 1] = tmp
	}
}

func main() {
	a := []int{100, 4, 200, 1, 3, 2}
	insertSort(a)
	fmt.Println(a)
}
