package main

import (
	"fmt"
)

func UNUSED(...interface{}) {}

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

func main() {
	i := Node{"i", nil, nil}
	h := Node{"h", nil, nil}
	g := Node{"g", &h, &i}
	f := Node{"f", nil, nil}
	e := Node{"e", nil, &g}
	d := Node{"d", nil, nil}
	c := Node{"c", nil, &f}
	b := Node{"b", &d, &e}
	a := Node{"a", &b, &c}

	fmt.Printf("depth: %v\n", depth(&a))

	fmt.Printf("breadth: %v\n", breadth(&a))
}

func depth(root *Node) []string {
	result := make([]string, 0)
	if root == nil {
		return result
	}
	result = append(result, (*root).Val)

	leftResult := depth((*root).Left)
	result = append(result, leftResult...)

	rightResult := depth((*root).Right)
	result = append(result, rightResult...)

	return result
}

func breadth(root *Node) []string {
	queue := make([]Node, 0)
	result := make([]string, 0)

	if root == nil {
		return result
	}

	queue = append(queue, *root)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		result = append(result, current.Val)

		if current.Left != nil {
			queue = append(queue, *(current.Left))
		}
		if current.Right != nil {
			queue = append(queue, *(current.Right))
		}
	}

	return result
}

func reverse(root *Node) *Node {
	if root == nil { return nil }

	left := reverse(root.Left)
	right := reverse(root.Right)

	root.Left = right
	root.Right = left

	return root
}
