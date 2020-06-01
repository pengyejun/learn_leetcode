package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := candies[0]
	for _, v := range candies[1:] {
		if max < v {
			max = v
		}
	}
	max -= extraCandies
	returned := make([]bool, len(candies), len(candies))
	for i, v := range candies {
		returned[i] = v >= max
	}
	return returned
}

func main() {
	kidsWithCandies([]int{2,3,5,1,3}, 3)
}
