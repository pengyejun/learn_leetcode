package main

import "math"

func maxProduct(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	max := math.MinInt64
	imin := 1
	imax := 1
	for _, num := range nums {
		if num < 0 {
			imin, imax = imax, imin
		}
		tmp_imin := imin * num
		tmp_imax := imax * num
		if tmp_imax > num {
			imax = tmp_imax
		} else {
			imax = num
		}
		if tmp_imin < num {
			imin = tmp_imin
		} else {
			imin = num
		}
		if max < imax {
			max = imax
		}
	}
	return max
}


func main() {
	
}
