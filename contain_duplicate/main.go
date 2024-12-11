package main

import "fmt"

func checkDuplicate(a []int) bool {
	var duplicate = make(map[int]bool, len(a))
	for _, value := range a {
		if _, ok := duplicate[value]; ok {
			return true
		} else {
			duplicate[value] = true
		}
	}
	return false
}

func main() {
	var a = []int{1, 2, 3, 3}
	fmt.Printf("Expected Output: %+v\n", checkDuplicate(a))

	var b = []int{1, 2, 3, 4}
	fmt.Printf("Expected Output: %+v\n", checkDuplicate(b))
}
