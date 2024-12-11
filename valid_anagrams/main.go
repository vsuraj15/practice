package main

import (
	"fmt"
	"reflect"
)

func checkValidAnagram(a, b string) bool {
	var map1 = make(map[string]int, 0)
	var map2 = make(map[string]int, 0)
	for _, ch := range a[0:] {
		val := string(ch)
		if data, dataOk := map1[val]; dataOk {
			map1[val] = data + 1
		} else {
			map1[val] = 1
		}
	}
	for _, ch := range b[0:] {
		val := string(ch)
		if data, dataOk := map2[val]; dataOk {
			map2[val] = data + 1
		} else {
			map2[val] = 1
		}
	}
	if reflect.DeepEqual(map1, map2) {
		return true
	}
	return false
}

func main() {
	var s = "racecar"
	var t = "carrace"
	fmt.Printf("Expected Output: %+v\n", checkValidAnagram(s, t))

	var s1 = "jar"
	var t1 = "jam"
	fmt.Printf("Expected Output: %+v\n", checkValidAnagram(s1, t1))
}
