package main

import (
	"fmt"
	"strings"
)

func longestSubstring(str string) int {
	for i, char := range str {
		if strings.IndexRune(str, char) == len(str) {
			return 1
		}
		if !strings.Contains(str[i+1:], string(char)) {
			return len(str[i:])
		}
	}
	return 0
}

func main() {
	var str string
	str = "zxyzxyz"
	fmt.Printf("Exxpected longest string length output: %+v\n", longestSubstring(str))

	str = "xxxx"
	fmt.Printf("Exxpected longest string length output: %+v\n", longestSubstring(str))
}
