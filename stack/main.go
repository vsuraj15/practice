package main

import "fmt"

type Stack struct {
	item []int
}

func (s *Stack) Push(item int) {
	s.item = append(s.item, item)
}

func (s *Stack) Pop() int {
	if len(s.item) == 0 {
		return -1
	}
	item, items := s.item[len(s.item)-1], s.item[0:len(s.item)-1]
	s.item = items
	return item
}

func main() {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	fmt.Println("Poped Item: ", stack.Pop())
	fmt.Println("Poped Item: ", stack.Pop())
	fmt.Println("Poped Item: ", stack.Pop())
	fmt.Println("Poped Item: ", stack.Pop())
	fmt.Println("Poped Item: ", stack.Pop())
}
