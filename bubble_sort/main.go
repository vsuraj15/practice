package main

import "fmt"

func bubbleSort(data []int) []int {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	return data
}

func main() {
	data := []int{11, 14, 3, 8, 18, 17, 43}
	fmt.Printf("Actual Data: %+v\n", data)
	fmt.Printf("Sorted Data: %+v\n", bubbleSort(data))
}
