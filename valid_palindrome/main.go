package main

import (
	"fmt"
	"strings"
)

func reverseStr(str string) string {
	char := []rune(str)
	for i, j := 0, len(str)-1; i < len(char)/2; i, j = i+1, j-1 {
		char[i], char[j] = char[j], char[i]
	}
	return string(char)
}

func validPalindrom(str string) bool {
	str = strings.ReplaceAll(str, " ", "")
	if str == reverseStr(str) {
		return true
	}
	return false
}

func main() {
	var input string
	input = "Was it a car or a cat I saw"
	fmt.Printf("Expected output: %+v\n", validPalindrom(strings.ToLower(input)))

	input = "tab a cat"
	fmt.Printf("Expected output: %+v\n", validPalindrom(strings.ToLower(input)))
}
