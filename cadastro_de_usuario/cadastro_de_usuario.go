package main

import "fmt"

type Cadastro struct {
	cpf   int
	nome  string
	idade int
}

func main() {
	var cadastros [2]Cadastro

	for i := 0; i < 2; i++ {
		fmt.Println("Informe o CPF(apenas digitos): ")
		fmt.Scanln(&cadastros[i].cpf)
		fmt.Println("Informe o nome: ")
		fmt.Scanln(&cadastros[i].nome)
		fmt.Println("Informe a idade: ")
		fmt.Scanln(&cadastros[i].idade)
	}
	for i := 0; i < 2; i++ {
		fmt.Printf("O nome do Usuario %d é %s\n", i+1, cadastros[i].nome)
	}
}
