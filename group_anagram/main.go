package main

import (
	"fmt"
	"sort"
	"strconv"
)

func groupAnagram(strs []string) [][]string {
	var result = make([][]string, 0)
	if len(strs) == 0 {
		return result
	}
	var dic = map[string][]string{}
	for _, str := range strs {
		chars := []byte(str)
		ints := make([]int, len(chars))
		for _, char := range chars {
			ints = append(ints, int(char))
		}
		sort.Ints(ints)
		key := ""
		for _, i := range ints {
			key += strconv.Itoa(i)
		}
		if _, ok := dic[key]; !ok {
			dic[key] = make([]string, 0)
		}
		dic[key] = append(dic[key], str)
	}
	for _, res := range dic {
		result = append(result, res)
	}
	return result
}

func main() {
	var input []string
	input = []string{"act", "pots", "tops", "cat", "stop", "hat"}
	fmt.Printf("Expected Output: %+v\n", groupAnagram(input))

	input = []string{"x"}
	fmt.Printf("Expected Output: %+v\n", groupAnagram(input))

	input = []string{""}
	fmt.Printf("Expected Output: %+v\n", groupAnagram(input))
}
