package main

import "fmt"

func main() {
	input := 12321
	fmt.Println(isPalindrome(input))
	input = 56
	fmt.Println(isPalindrome(input))
}

func isPalindrome(n int) bool {
	input := n
	if input < 0 || (input != 0 && input%10 == 0) {
		return false
	}
	ss := 0
	for {
		q := input / 10
		r := input % 10
		ss = ss*10 + r
		input = q
		if q == 0 {
			break
		}
	}
	if ss == n {
		return true
	}
	return false
}
