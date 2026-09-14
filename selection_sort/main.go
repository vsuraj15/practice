package main

import "fmt"

func selectionSort(data []int) []int {
	for i := 0; i < len(data)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[minIndex] {
				minIndex = j
			}
		}
		temp := data[i]
		data[i] = data[minIndex]
		data[minIndex] = temp
	}
	return data
}

func main() {
	data := []int{2, 4, 3, 1, 6, 8, 5}
	fmt.Printf("Given Data: %+v\n", data)
	fmt.Printf("Sorted Data: %+v\n", selectionSort(data))
}
