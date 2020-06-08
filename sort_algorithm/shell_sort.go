package main

import "fmt"

func shellSort(nums []int) {
	length := len(nums)

	d := length / 2
	for d > 0 {
		for i := d; i < length; i++ {
			j := i
			for j >= d && nums[j - d] > nums[j] {
				nums[j - d], nums[j] = nums[j], nums[j - d]
				j -= d
			}
		}
		d = d / 2
	}
}

func main() {
	//a := []int{100, 4, 200, 1, 3, 2}
	a := []int{12, 1, 9, 21, 3, 20, 0}
	shellSort(a)
	fmt.Println(a)
}
