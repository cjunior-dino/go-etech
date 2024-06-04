package main

import "fmt"

type Node struct {
	key   int
	left  *Node
	right *Node
}

func (n *Node) Insert(c int) {
	if n.key < c {
		if n.left == nil {
			n.left = &Node{key: c}
		} else {
			n.left.Insert(c)
		}
	} else if n.key > c {
		if n.right == nil {
			n.right = &Node{key: c}
		} else {
			n.right.Insert(c)
		}
	}
}

func printInDorder(i *Node) {
	if i != nil {
		printInDorder(i.right)
		fmt.Println(i.key)
		printInDorder(i.left)
	}
}

func printInOrder(i *Node) {
	if i != nil {
		printInOrder(i.left)
		fmt.Println(i.key)
		printInOrder(i.right)
	}

}
func main() {
	tree := &Node{}
	var x, i int

	for {
		fmt.Println("Escolha uma opção:")
		fmt.Println("1- inserir dado:")
		fmt.Println("2- mostrar em ordem crescente:")
		fmt.Println("3- mostrar em ordem decrescente:")
		fmt.Scanln(&x)

		if x == 1 {
			fmt.Scanln(&i)
			tree.Insert(i)
		} else if x == 2 {
			printInOrder(tree)
		} else if x == 3 {
			printInDorder(tree)
		}

	}

}
