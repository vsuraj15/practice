package main

import "fmt"

type List struct {
	head *Node
	tail *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Push(val int) {
	node := &Node{value: val}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
	node.next = l.head
}

func Traverse(n *List) {
	node := n.First()
	nodeIndex := 1
	for {
		fmt.Printf("Node %d: %+v\n", nodeIndex, node)
		fmt.Printf("Value of next node: %d\n", node.next.value)
		node = node.Next()
		nodeIndex++
		if node == n.First() {
			break
		}
	}
}

type Node struct {
	value int
	next  *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func main() {
	l := &List{}
	l.Push(100)
	l.Push(200)
	l.Push(300)
	l.Push(400)
	head := l.First()
	fmt.Printf("First Node: %+v\n", head)
	fmt.Println("----------- Traverse each node ------------")
	Traverse(l)
}
