package main

import (
	"fmt"
	"math"
)

func buySellStock(input []int) int {
	var profit int
	minimum := math.MaxInt32
	for _, num := range input {
		if num < minimum {
			minimum = num
		}

		if num-minimum > 0 {
			profit += num - minimum
			minimum = num
		}
	}
	return profit
}

func main() {
	var input []int
	input = []int{10, 1, 5, 6, 7, 1}
	fmt.Printf("Expected profit output: %+v\n", buySellStock(input))

	input = []int{10, 8, 7, 5, 2}
	fmt.Printf("Expected profit output: %+v\n", buySellStock(input))
}
