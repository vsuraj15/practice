package main

import "fmt"

func frequentElement(arr []int, count int) []int {
	var arrMap = make(map[int]int, 0)
	for _, value := range arr {
		if val, ok := arrMap[value]; ok {
			arrMap[value] = val + 1
		} else {
			arrMap[value] = 1
		}
	}
	var result = make([]int, count)
	pos := 0
	for key, counter := range arrMap {
		if counter > 1 {
			result[pos] = key
			pos++
		}
	}
	return result
}

func main() {
	var input []int
	var count int
	input = []int{1, 2, 2, 3, 3, 3}
	count = 2
	fmt.Printf("Expected Output: %+v\n", frequentElement(input, count))

	input = []int{7, 7}
	count = 1
	fmt.Printf("Expected Output: %+v\n", frequentElement(input, count))
}
