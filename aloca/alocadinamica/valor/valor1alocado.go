package main

import (
	"fmt"
	"unsafe"
)

func main() {
	num := new(int)
	*num = 1

	fmt.Println("Valor alocado:", *num)
	fmt.Println("Espaço usado:", unsafe.Sizeof(*num), "bytes")

	var resposta string
	fmt.Println("Voce quer liberar a memoria? (s/n)")
	fmt.Scanln(&resposta)

	if resposta == "s" {
		num = nil
		fmt.Println("Memoria liberada com sucesso!")
		if num == nil {
			fmt.Println("Nao há nenhum valor alocado.")
		}
	} else {
		fmt.Println("Memoria não liberada.")
	}
}
