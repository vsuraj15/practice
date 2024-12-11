package main

import "fmt"

func twoSum(a []int, target int) []int {
	var res = make(map[int]int, 0)
	for i, j := range a {
		complement := target - j
		if num, ok := res[complement]; ok {
			return []int{num, i}
		}
		res[j] = i
	}
	return []int{}
}

func main() {
	a1 := []int{3, 4, 5, 6}
	target1 := 7
	fmt.Printf("Expected Output: %+v\n", twoSum(a1, target1))

	a2 := []int{4, 5, 6}
	target2 := 10
	fmt.Printf("Expected Output: %+v\n", twoSum(a2, target2))

	a3 := []int{5, 5}
	target3 := 10
	fmt.Printf("Expected Output: %+v\n", twoSum(a3, target3))
}
