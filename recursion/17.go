package main

var m = map[rune]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}

func letterCombinations(digits string) []string {
	length := len(digits)
	if length == 0 {
		return nil
	}
	res1 := letterCombinations(digits[1:])
	if res1 == nil {
		res1 = []string{""}
	}
	res := make([]string, 0)
	for _, v := range m[rune(digits[0])] {
		for _, v1 := range res1 {
			res = append(res, string(v) + v1)
		}
	}
	return res
}

func main() {
	
}
