package main

import "fmt"

type Queue struct {
	item chan int
}

func (q *Queue) Enqueue(item int) {
	q.item <- item
}

func (q *Queue) Dequeue() int {
	return <-q.item
}

func main() {
	stack := Queue{item: make(chan int, 4)}
	stack.Enqueue(1)
	stack.Enqueue(2)
	stack.Enqueue(3)
	stack.Enqueue(4)
	fmt.Println("Dequeued Item: ", stack.Dequeue())
	fmt.Println("Dequeued Item: ", stack.Dequeue())
	fmt.Println("Dequeued Item: ", stack.Dequeue())
	fmt.Println("Dequeued Item: ", stack.Dequeue())
}
