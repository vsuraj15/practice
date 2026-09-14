package main

import "fmt"

func main() {
	input := 5
	fmt.Println(reverseInteger(input))
	input = 15
	fmt.Println(reverseInteger(input))
}

func reverseInteger(n int) int {
	input := n
	divisor := 1
	newInt := 0
	if input < 10 {
		return input
	}
	for {
		if divisor == 0 {
			break
		}
		divisor = input / 10
		remainder := input % 10
		newInt = newInt*10 + remainder
		input = divisor
	}
	return newInt
}
