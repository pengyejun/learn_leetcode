package main

import (
	"fmt"
	"unsafe"
)

func addBinary(a string, b string) string {
	carry := uint8(0)
	res := make([]byte, 0)
	for i := 0; i < len(a) || i < len(b); i++ {
		sum := carry
		if i < len(a) {
			sum += a[len(a)- i - 1] - 48
		}
		if i < len(b) {
			sum += b[len(b) - i - 1] - 48
		}
		if sum >= 2 {
			carry = 1
			sum -= 2
		} else {
			carry = 0
		}
		res = append([]byte{sum + 48}, res...)
	}
	if carry != 0 {
		res = append([]byte{49}, res...)
	}
	return *(*string)(unsafe.Pointer(&res))
}


func main() {
	fmt.Println(addBinary("11", "1"))
}
