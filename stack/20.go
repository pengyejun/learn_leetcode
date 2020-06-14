package main

var m = map[rune]rune{'(': ')', '{': '}', '[': ']'}
func isValid(s string) bool {
	length := len(s)
	if length == 0 {
		return true
	}
	if length % 2 != 0 {
		return false
	}
	stack := make([]rune, 0)
	for i := 0; i < length; i++ {
		v := rune(s[i])
		if _, ok := m[v]; ok {
			stack = append(stack, v)
		}else {
			if len(stack) <= 0 {
				return false
			}
			if v != m[stack[len(stack) - 1]] {
				return false
			}
			stack = stack[:len(stack) - 1]
		}
	}
	return len(stack) == 0
}

func main() {
	
}
