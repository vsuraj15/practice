package main

import "fmt"

func Merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	i, j := 0, 0
	for k := 0; k < len(result); k++ {
		if i >= len(left) {
			result[k] = right[j]
			j++
		} else if j >= len(right) {
			result[k] = left[i]
			i++
		} else if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
	}
	return result
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	done := make(chan bool)
	var left []int
	go func() {
		left = MergeSort(arr[:mid])
		done <- true
	}()
	right := MergeSort(arr[mid:])
	<-done
	return Merge(left, right)

}

func main() {
	data := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}
	fmt.Printf("Actual Array: %+v\n", data)
	fmt.Printf("Sorted Array: %+v\n", MergeSort(data))
}
