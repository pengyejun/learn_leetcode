package main

import "fmt"

func quickSort(nums []int, l, r int) {
	if l < r {
		i := l
		j := r
		x := nums[l]
		for i < j {
			for i < j && nums[j] >= x {
				j--
			}
			if i < j {
				nums[i] = nums[j]
				i++
			}
			for i < j && nums[i] < x {
				i++
			}
			if i < j {
				nums[j] = nums[i]
				j--
			}
		}
		nums[i] = x
		quickSort(nums, l, i - 1)
		quickSort(nums, i + 1, r)
	}
}

func main() {
	a := []int{12, 1, 9, 21, 3, 20, 0}
	quickSort(a, 0, len(a) - 1)
	fmt.Println(a)
}
