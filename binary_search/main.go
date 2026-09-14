package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	target := 3
	if found := binarySearch(data, target); found {
		fmt.Println("Found expected target")
	} else {
		fmt.Println("Missing value")
	}
}

func binarySearch(data []int, target int) bool {
	low, high := 0, len(data)-1
	for low <= high {
		mid := (low + high) / 2
		if data[mid] == target {
			return true
		} else if data[mid] < target {
			low = mid + 1
		} else {
			high = mid + 1
		}
	}
	return false
}
