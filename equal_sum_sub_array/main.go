package main

import "fmt"

func main() {
	var list []int
	list = []int{1}
	fmt.Printf("Expected output be: %+v, For List: %+v\n", equalSubArray(list), list)

	list = []int{1, 2, 2, 3}
	fmt.Printf("Expected output be: %+v, For List: %+v\n", equalSubArray(list), list)

	list = []int{1, 2, 3}
	fmt.Printf("Expected output be: %+v, For List: %+v\n", equalSubArray(list), list)

	list = []int{2, 3, 5}
	fmt.Printf("Expected output be: %+v, For List: %+v\n", equalSubArray(list), list)

	list = []int{1, 7, 3, 5}
	fmt.Printf("Expected output be: %+v, For List: %+v\n", equalSubArray(list), list)

	list = []int{-4, 1, 1, 1, 1}
	fmt.Printf("Expected output be: %+v, For List: %+v\n", equalSubArray(list), list)

	list = []int{4, 1, 1, 1, 1}
	fmt.Printf("Expected output be: %+v, For List: %+v\n", equalSubArray(list), list)

	list = []int{1, 1, 1, 1, 4}
	fmt.Printf("Expected output be: %+v, For List: %+v\n", equalSubArray(list), list)

}

func equalSubArray(list []int) [][]int {
	output := make([][]int, 0)
	if len(list) < 2 {
		return output
	}
	splitIndex := getSplitPoint(list)
	if splitIndex == -1 || splitIndex == len(list) {
		return output
	}
	output = append(output, list[:splitIndex])
	output = append(output, list[splitIndex:])
	return output
}

func getSplitPoint(list []int) int {
	leftSum := 0
	for i := range len(list) {
		leftSum += list[i]
		rightSum := 0
		for j := i + 1; j < len(list); j++ {
			rightSum += list[j]
		}
		if leftSum == rightSum {
			return i + 1
		}
	}
	return -1
}
