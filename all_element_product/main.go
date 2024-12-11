package main

import "fmt"

func main() {
	list := []int{5, 3, 7, 2, 1}
	fmt.Printf("Expected Output: %+v\n", productOfAllElements(list))

	list = []int{2, 4, 5, 2, 3}
	fmt.Printf("Expected Output: %+v\n", productOfAllElements(list))
}

func productOfAllElements(list []int) []int {
	if len(list) == 0 {
		return list
	}
	arr := make([]int, len(list))
	product := 1
	for i := len(list) - 1; i >= 0; i-- {
		product *= list[i]
		arr[i] = product
	}
	product = list[0]
	for i := 0; i < len(list); i++ {
		if i == 0 {
			list[i] = arr[1]
			continue
		}
		value := list[i]
		if i != len(list)-1 {
			list[i] = arr[i+1] * product
		} else {
			list[i] = product
		}
		product *= value
	}
	return list
}
