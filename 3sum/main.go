package main

import (
	"fmt"
	"sort"
)

func threeSum(ints []int) [][]int {
	var ansList = make([][]int, 0)
	length := len(ints)
	sort.Ints(ints)
	for i := 0; i < length-2; i++ {
		low, high := i+1, length-1
		sum := 0 - ints[i]
		for {
			if low >= high {
				break
			}
			switch {
			case ints[low]+ints[high] == sum:
				var ans = []int{ints[i], ints[low], ints[high]}
				ansList = append(ansList, ans)
				low++
				high = high - 1
			case ints[low]+ints[high] < sum:
				low += 1
			case ints[low]+ints[high] > sum:
				high = high - 1
			}
		}
	}
	return ansList
}

func main() {
	var input []int
	input = []int{-1, 0, 1, 2, -1, -4}
	fmt.Printf("Expected Output: %+v\n", threeSum(input))
}
