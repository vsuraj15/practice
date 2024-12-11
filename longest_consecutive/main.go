package main

import (
	"fmt"
	"sort"
)

func longestConsecutiveSequence(ints []int) int {
	counter := 1
	sort.Ints(ints)
	for i := 0; i < len(ints)-1; i++ {
		if ints[i]+1 == ints[i+1] {
			counter++
		}
	}
	return counter
}

func main() {
	var input []int
	input = []int{2, 20, 4, 10, 3, 4, 5}
	fmt.Printf("Expected Output: %+v\n", longestConsecutiveSequence(input))

	input = []int{0, 3, 2, 5, 4, 6, 1, 1}
	fmt.Printf("Expected Output: %+v\n", longestConsecutiveSequence(input))
}
