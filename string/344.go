package main

func reverseString(s []byte)  {
	length := len(s) - 1
	if length < 1 {
		return
	}
	mid := length / 2
	for i := 0; i <= mid; i++ {
		s[i], s[length - i] = s[length - i], s[i]
	}
}

func main() {
	
}
