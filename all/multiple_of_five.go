package main

import "fmt"

func main() {
	input := 15
	fmt.Println(solve(input))

	input = 7
	fmt.Println(solve(input))
}

func solve(input int) []int {
	result := make([]int, 0)
	for i := 1; i <= input; i++ {
		if i%5 == 0 {
			result = append(result, i)
		}
	}
	return result
}
