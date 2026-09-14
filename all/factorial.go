package main

import "fmt"

func main() {
	input := 5
	fmt.Println(factorial(input))
}

func factorial(input int) int {
	output := 1
	for i := input; i >= 1; i-- {
		output *= i
	}
	return output
}
