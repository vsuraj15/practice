package main

import (
	"fmt"
	"strconv"
)

func main() {
	var list []int
	list = []int{0}
	fmt.Printf("Expected output: %+v, list: %+v\n", getDuplicateItem(list), list)

	list = []int{1, 2, 2, 3}
	fmt.Printf("Expected output: %+v, list: %+v\n", getDuplicateItem(list), list)

	list = []int{}
	fmt.Printf("Expected output: %+v, list: %+v\n", getDuplicateItem(list), list)

	list = []int{1, 2, 4, 5, 5, 5}
	fmt.Printf("Expected output: %+v, list: %+v\n", getDuplicateItem(list), list)

	list = []int{3, 2, 1, 4, 5, 4}
	fmt.Printf("Expected output: %+v, list: %+v\n", getDuplicateItem(list), list)
}

func getDuplicateItem(arr []int) string {
	var duplicate = make(map[int]bool)
	for _, value := range arr {
		if duplicate[value] {
			return strconv.Itoa(value)
		}
		duplicate[value] = true
	}
	return "NA"
}
