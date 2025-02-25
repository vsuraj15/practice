package main

import "fmt"

type List struct {
	head *Node
	tail *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) Push(val int) {
	node := &Node{value: val}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
	}
	l.tail = node
}

func (l *List) Find(value int) (*Node, error) {
	node := l.First()
	for {
		if node == nil {
			return &Node{}, fmt.Errorf("Node not found for value: %d", value)
		}
		if node.value == value {
			return node, nil
		}
		node = node.Next()
	}
}

func TraverseForward(n *List) {
	node := n.First()
	nodeIndex := 1
	for {
		fmt.Printf("Node %d: %+v\n", nodeIndex, node)
		node = node.Next()
		nodeIndex++
		if node == nil {
			break
		}
	}
}

func TraverseBackward(n *List) {
	node := n.Last()
	nodeIndex := 1
	for {
		fmt.Printf("Node %d: %+v\n", nodeIndex, node)
		node = node.Prev()
		nodeIndex++
		if node == nil {
			break
		}
	}
}

type Node struct {
	value int
	next  *Node
	prev  *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func main() {
	l := &List{}
	l.Push(100)
	l.Push(200)
	l.Push(300)
	l.Push(400)
	fmt.Println("----------- Traverse forward each node ------------")
	TraverseForward(l)
	fmt.Println("----------- Traverse backward each node ------------")
	TraverseBackward(l)
	// Find an valid element
	validElement, validErr := l.Find(300)
	if validErr != nil {
		fmt.Println("Error while find a value in list.Err: %+v\n", validErr)
	} else {
		fmt.Printf("Node Found: %+v\n", validElement)
	}
	// Find an invalid element
	invalidElement, invalidErr := l.Find(600)
	if invalidErr != nil {
		fmt.Printf("Error while find a value in list.Err: %+v\n", invalidErr)
	} else {
		fmt.Printf("Node Found: %+v\n", invalidElement)
	}
}
