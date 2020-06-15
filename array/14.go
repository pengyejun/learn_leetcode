package main

import "unsafe"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	if len(strs) == 0 {
		return ""
	}
	res := make([]byte, 0)
	for i, v := range strs[0] {
		v := uint8(v)
		for _, str := range strs[1:] {
			if len(str) <= i || str[i] != v {
				return *(*string)(unsafe.Pointer(&res))
			}
		}
		res = append(res, v)
	}
	return strs[0]
}

func main() {
	longestCommonPrefix([]string{"flower","flow","flight"})
}
