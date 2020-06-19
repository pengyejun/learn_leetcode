package main

import "fmt"

func isPalindrome(s string) bool {
	length := len(s)
	if length <= 1 {
		return true
	}

	l, r := 0, length - 1
	for l < r {
		// 校验字符是否符合0-9 a-z A-Z
		if s[l] <= 47 || (s[l] >= 58 && s[l] <= 64) || (s[l] <= 96 && s[l] >= 91) || s[l] >= 123 {
			l++
			continue
		}
		if s[r] <= 47 || (s[r] >= 58 && s[r] <= 64) || (s[r] <= 96 && s[r] >= 91) || s[r] >= 123 {
			r--
			continue
		}
		tmpL := s[l]
		tmpR := s[r]
		if 65 <= tmpL && tmpL <= 90 {
			tmpL += 32
		}
		if 65 <= tmpR && tmpR <= 90 {
			tmpR += 32
		}
		if tmpL != tmpR {
			return false
		}
		l++
		r--
	}
	return true
}


func main() {
	fmt.Println(isPalindrome("race a car"))
}
