package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(k int) {
	if n.Key < k {
		fmt.Println(n.Key)
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		fmt.Println(n.Key)
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

func printInOrder(n *Node) {
	if n != nil {
		fmt.Println(n.Key)
	}
}

func main() {
	tree := &Node{Key: 10}
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(7)

	printInOrder(tree)

}
