package main


func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	root := 0
	for root != slow {
		root = nums[root]
		slow = nums[slow]
	}
	return slow
}

func findDuplicateMid(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		count := 0
		for _, v := range nums {
			if v <= mid {
				count++
			}
		}
		if count > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}


func main() {
	
}
