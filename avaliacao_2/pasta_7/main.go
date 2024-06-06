package main

import "fmt"

type Node struct { // estrutura do nó da BTS
	Key   int
	Left  *Node
	Right *Node
}

type BinarySearchtTree struct { // estrutura da BTS
	Root *Node
}

func (t *BinarySearchtTree) insert(key int) *BinarySearchtTree { // insere um novo nó na BTS
	if t.Root == nil {
		t.Root = &Node{Key: key}
	} else {
		t.Root.insert(key)
	}
	return t
}

func (n *Node) insert(key int) { // insere um dado na BTS
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

func (n *Node) Search(i int) bool { // realiza a busca de um nó na arvore
	if n == nil {
		return false
	} else if n.Key == i {
		return true
	} else if i < n.Key {
		return n.Left.Search(i)
	} else {
		return n.Right.Search(i)
	}

}

func (n *Node) Printarvore() { // exibe a arvore em ordem crescente
	if n == nil {
		return
	}

	n.Left.Printarvore()
	fmt.Println(" ")
	fmt.Println(n.Key)
	fmt.Println(" ")
	n.Right.Printarvore()
}

func main() {
	registro := 1

	tree := &BinarySearchtTree{}
	var x int
	for {
		fmt.Println("Escolha uma das opções")
		fmt.Println(" ")
		fmt.Println("1- Cadastrar um aluno:")
		fmt.Println("2- Procurar se aluno existe:")
		fmt.Println("3- Exibir todos os cadastros:")
		fmt.Println("4- Sair do sistema:")
		fmt.Scanln(&x)

		if x == 1 {
			tree.insert(registro)
			fmt.Printf("Parabens o aluno %d esta cadastrado no registro escolar !!!!", registro)
			fmt.Println(" ")
			registro++
		} else if x == 2 {
			var procura int
			fmt.Println("Informe o codigo do aluno:")
			fmt.Scanln(&procura)
			if tree.Root.Search(procura) {
				fmt.Println(" ")
				fmt.Println("O aluno existe no registro!!")
			} else {
				fmt.Println(" ")
				fmt.Println("O aluno não existe no registro!!")
			}
		} else if x == 3 {
			tree.Root.Printarvore()
		} else if x == 4 {
			fmt.Println(" ")
			fmt.Println("Sistema encerrado !!!")
			break
		} else {
			fmt.Println(" ")
			fmt.Println("Opção invalida")
			fmt.Println(" ")
		}
	}

}

/*
                   __
                  /°_)
           .-^^^-/ /
        __/       /
       <__.|_|-|_|
   By: DinoKaminari
*/
