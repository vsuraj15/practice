package main

import (
	"fmt"
	"strings"
)

func encodeString(strs []string) string {
	return strings.Join(strs, ":")
}

func decodeString(str string) []string {
	str = strings.ReplaceAll(str, "::", ":")
	var arr = make([]string, 0)
	arr = strings.Split(str, ":")
	for i, value := range arr {
		if value == "" {
			arr[i] = ":"
		}
	}
	return arr
}

func main() {
	var input []string
	input = []string{"neet", "code", "love", "you"}
	fmt.Printf("Expected Decoded Output: %s\n", decodeString(encodeString(input)))

	input = []string{"we", "say", ":", "yes"}
	fmt.Printf("Expected Decoded Output: %s\n", decodeString(encodeString(input)))
}
