package main

import "fmt"

func main() {
	list := []int{7, 5, 3, 9, 8, 6, 2}
	startIndex, endIndex := 2, 6
	fmt.Printf("Expected output: %+v\n", reverseList(list, startIndex, endIndex))

	list = []int{5, 0, 3, 7, 6, 2, 5, 9, 8}
	startIndex, endIndex = 1, 7
	fmt.Printf("Expected output: %+v\n", reverseList(list, startIndex, endIndex))
}

func reverseList(list []int, start, end int) []int {
	if start > end {
		return list
	}
	for i := start; i <= (start+end)/2 && i < (end-i+start); i++ {
		list[i], list[end-i+start] = list[end-i+start], list[i]
	}
	return list
}
