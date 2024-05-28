package main

import (
	"fmt"
	"unsafe"
)

func main() {
	data := make([]int, 1e6)
	fmt.Println("Valor alocado: ", data[0])
	fmt.Println("espaço usado: ", unsafe.Sizeof(data))

	var reposta string
	fmt.Println("Voce quer liberar a memoria? (s/n)")
	fmt.Scanln(&reposta)

	if reposta == "s" {
		data = nil
		fmt.Println("Memoria liberada")
	} else {
		fmt.Println("Memoria não librada")
	}

}
