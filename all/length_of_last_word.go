package main

import "fmt"

func main() {
	s := "Hello World"
	fmt.Println("Length of last word: ", lengthOfLastWord(s))
}

func lengthOfLastWord(input string) int {
	var length int
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] != ' ' {
			length++
		} else {
			return length
		}
	}
	return length
}
