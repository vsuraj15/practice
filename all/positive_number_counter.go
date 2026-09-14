package main

import "fmt"

func main() {
	input := []int{-2, 5, 7, -1, 3}
	fmt.Println(counter(input))
	input = []int{-2, -10}
	fmt.Println(counter(input))
}

func counter(input []int) int {
	count := 0
	for _, value := range input {
		if value > 0 {
			count++
		}
	}
	return count
}
