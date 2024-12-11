package main

import "fmt"

type Tree struct {
	node *Node
}

func (t *Tree) insert(val int) *Tree {
	if t.node == nil {
		t.node = &Node{value: val}
	} else {
		t.node.insert(val)
	}
	return t
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) isExists(value int) bool {
	if n == nil {
		return false
	}
	if value == n.value {
		return true
	}
	if value <= n.value {
		fmt.Println("Left node")
		return n.left.isExists(value)
	} else {
		fmt.Println("Right node")
		return n.right.isExists(value)
	}
}

func (n *Node) insert(value int) {
	if value <= n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			n.right.insert(value)
		}
	}
}

func printTree(n *Node) {
	if n == nil {
		return
	}
	fmt.Printf("%d\n", n.value)
	printTree(n.left)
	printTree(n.right)
}

func main() {
	t := &Tree{}
	t.insert(10).insert(8).insert(20).insert(9).insert(0).insert(15).insert(25)
	printTree(t.node)
	fmt.Println("IsExists: ", t.node.isExists(27))
}
