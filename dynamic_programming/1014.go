package main

func maxScoreSightseeingPair(A []int) int {
	all, left := 0, A[0]
	for i := 1; i < len(A); i++ {
		if left + A[i] - i > all {
			all = left + A[i] - i
		}
		if A[i] + i > left {
			left = A[i] + i
		}
	}
	return all
}


func main() {
	maxScoreSightseeingPair([]int{8, 1, 5, 2, 7})
}
