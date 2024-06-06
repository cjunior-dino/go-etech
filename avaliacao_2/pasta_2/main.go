package main

import (
	"fmt"
)

type Stack struct {
	data []int
}

func (s *Stack) Push(val int) { // insere um dado na pilha
	s.data = append(s.data, val)
}

func (s *Stack) Pop() int { // remove o topo da pilha
	if len(s.data) == 0 {
		return -1
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}

func (s *Stack) Peek() int { // visuzaliza o topo da pilha
	if len(s.data) == 0 {
		return -1
	}
	return s.data[len(s.data)-1]
}

func (s *Stack) IsEmpty() bool { // verifica se a pilha esta vazia
	return len(s.data) == 0
}

func main() {
	stack := Stack{}

	registro := 1
	var x int
	for {
		x = 0
		fmt.Println("BEM-VINDO AO HOSPILTAL SAUDE LEGAL")
		fmt.Println(" ")
		fmt.Println("Escolha uma das opções")
		fmt.Println("1- Adicionar um novo registro:")
		fmt.Println("2- Revisar o ultimo registro:")
		fmt.Println("3- Ver ultimo regitro da pilha:")
		fmt.Println("4- Verificar se ainda há registro:")
		fmt.Println("5- Sair do sistema:")
		fmt.Scanln(&x)

		if x == 1 {
			stack.Push(registro)
			fmt.Printf("Registro de numero %d cadastrado !!!!", registro)
			fmt.Println(" ")
			registro++
		} else if x == 2 {
			if stack.IsEmpty() {
				fmt.Println("Não existe registro")
				fmt.Println(" ")
			} else {
				fmt.Printf("Verificado o registro N° %d", stack.Pop())
				fmt.Println(" ")

			}
		} else if x == 3 {
			if stack.IsEmpty() {
				fmt.Println("Não existe registro")
				fmt.Println(" ")
			} else {
				fmt.Printf("O ultimo registro é %d", stack.Peek())
				fmt.Println(" ")
			}
		} else if x == 4 {
			if stack.IsEmpty() {
				fmt.Println("Não existe registro")
				fmt.Println(" ")
			} else {
				fmt.Println("Ainda existe registro para verificar")
				fmt.Println(" ")
			}
		} else if x == 5 {
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
