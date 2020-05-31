package main

func subarraysDivByK(A []int, K int) int {
	sum := 0
	ans := 0

	arr := make([]int, K, K)
	arr[0] = 1
	for i := 1; i <= len(A); i++ {
		sum = sum + A[i - 1]
		key := (sum % K + K) % K
		ans += arr[key]
		arr[key]++
	}
	return ans
}


func main() {
	subarraysDivByK([]int{4,5,0,-2,-3,1}, 5)
}
