package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

type BinarySearchtTree struct {
	Root *Node
}

func (t *BinarySearchtTree) insert(key int) *BinarySearchtTree {
	if t.Root == nil {
		t.Root = &Node{Key: key}
	} else {
		t.Root.insert(key)
	}
	return t
}

func (n *Node) insert(key int) {
	if n == nil {
		return
	} else if key <= n.Key {
		if n.Left == nil {
			n.Left = &Node{Key: key}
		} else {
			n.Left.insert(key)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Key: key}
		} else {
			n.Right.insert(key)
		}
	}
}

func (n *Node) imprime() {
	if n == nil {
		return
	}

	fmt.Println(n.Key)
	n.Left.imprime()
	n.Right.imprime()
}

func main() {
	tree := &BinarySearchtTree{}
	var x int
	for i := 0; i < 4; i++ {
		fmt.Scanln(&x)
		tree.insert(x)
	}
	fmt.Println(" ")
	tree.Root.imprime()

}
