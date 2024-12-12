package main

import (
	"fmt"
)

func containerWithMostWater(height []int) int {
	ans, l, r := 0, 0, len(height)-1
	for l < r {
		ans = max(ans, min(height[l], height[r])*(r-l))
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var height []int
	height = []int{1, 7, 2, 5, 4, 7, 3, 6}
	fmt.Printf("Expected Output: %v\n", containerWithMostWater(height))

	height = []int{2, 2, 2}
	fmt.Printf("Expected Output: %v\n", containerWithMostWater(height))
}
