package main

import "fmt"

func validParenthesis(str string) bool {
	p := []rune{}
	for _, c := range str {
		switch c {
		case '(':
			p = append(p, ')')
		case '[':
			p = append(p, ']')
		case '{':
			p = append(p, '}')
		default:
			if len(p) == 0 || p[len(p)-1] != c {
				return false
			}
			p = p[:len(p)-1]
		}
	}
	return len(p) == 0
}

func main() {
	var input string
	input = "[]"
	fmt.Printf("Expected output: %+v\n", validParenthesis(input))

	input = "([{}])"
	fmt.Printf("Expected output: %+v\n", validParenthesis(input))

	input = "[(])"
	fmt.Printf("Expected output: %+v\n", validParenthesis(input))
}
