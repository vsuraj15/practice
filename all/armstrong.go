package main

import (
	"fmt"
	"math"
)

func main() {
	input := 371
	fmt.Println(isArmstrongNumber(input))
	input = 71
	fmt.Println(isArmstrongNumber(input))
}

func isArmstrongNumber(n int) bool {
	input := n
	expectedOutput := n
	counter := 0
	diviser := 1
	for {
		if diviser == 0 {
			break
		}
		diviser = n / 10
		counter++
		n = diviser
	}
	count := counter
	sum := 0
	for {
		if counter == 0 {
			break
		}
		remainder := input % 10
		input = input / 10
		sum += int(math.Pow(float64(remainder), float64(count)))
		counter--
	}
	if sum == expectedOutput {
		return true
	}
	return false
}
