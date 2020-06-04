package main

func productExceptSelf(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	k := 1
	for i := 0; i < length; i++ {
		res[i] = k
		k *= nums[i]
	}

	k = 1
	for i := length - 1; i >= 0; i--{
		res[i] *= k
		k *= nums[i]
	}
	return res
}


func main() {
	
}
