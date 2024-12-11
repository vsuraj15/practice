package main

import (
	"fmt"
	"math"
)

func main() {
	num1 := []int{2, 9}
	num2 := []int{4, 9, 9}
	fmt.Printf("Addition of numbers: %+v, %v. Will give result: %+v\n", num1, num2, addNumbers(num1, num2))

	num1 = []int{1}
	num2 = []int{4, 9}
	fmt.Printf("Addition of numbers: %+v, %v. Will give result: %+v\n", num1, num2, addNumbers(num1, num2))

	num1 = []int{1}
	num2 = []int{}
	fmt.Printf("Addition of numbers: %+v, %v. Will give result: %+v\n", num1, num2, addNumbers(num1, num2))

	num1 = []int{1}
	num2 = []int{9, 9, 9}
	fmt.Printf("Addition of numbers: %+v, %v. Will give result: %+v\n", num1, num2, addNumbers(num1, num2))
}

func addNumbers(num1, num2 []int) []int {
	num1, num2 = equilizeLength(num1, num2)
	carry := false
	for i := len(num1) - 1; i > -1; i-- {
		num1[i] += num2[i]
		if carry {
			num1[i]++
			carry = false
		}
		if num1[i] >= 10 {
			num1[i] -= 10
			carry = true
		}
	}
	if carry {
		num1 = append([]int{1}, num1...)
	}
	return num1
}

func equilizeLength(num1, num2 []int) ([]int, []int) {
	diff := math.Abs(float64(len(num2) - len(num1)))
	zeros := make([]int, int(diff))
	if len(num2) > len(num1) {
		num1 = append(zeros, num1...)
	} else if len(num1) > len(num2) {
		num2 = append(zeros, num2...)
	}
	return num1, num2
}
