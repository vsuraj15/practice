package main

import (
	"fmt"
)

func productOfArr(ints []int) []int {
	var result = make([]int, len(ints))
	for index := 0; index < len(ints); index++ {
		var product = 1
		for i, num := range ints {
			if i != index {
				product = product * num
			}
		}
		result[index] = product
	}
	return result
}

func main() {
	var input []int
	input = []int{1, 2, 4, 6}
	fmt.Printf("Expected Output: %+v\n", productOfArr(input))
	input = []int{-1, 0, 1, 2, 3}
	fmt.Printf("Expected Output: %+v\n", productOfArr(input))
}
