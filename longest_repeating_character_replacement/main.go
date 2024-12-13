package main

import "fmt"

func longestRepeatingCharacterReplacement(input string, count int) int {
	maxCount := 0
	start := 0
	maxLength := 0
	countMap := make(map[byte]int, 0)

	for pos := 0; pos < len(input); pos++ {
		countMap[input[pos]]++
		maxCount = max(maxCount, countMap[input[pos]])
		if pos-start+1-maxCount > count {
			countMap[input[start]]--
			start++

		}
		maxLength = max(maxLength, pos-start+1)
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var input string
	var count int
	input = "XYYX"
	count = 2
	fmt.Printf("Expected output: %+v\n", longestRepeatingCharacterReplacement(input, count))

	input = "AAABABB"
	count = 1
	fmt.Printf("Expected output: %+v\n", longestRepeatingCharacterReplacement(input, count))
}
