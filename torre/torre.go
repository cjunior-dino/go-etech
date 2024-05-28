package main

import "fmt"

func hanoi(n int, origem, destino, aux string) {
	if n == 1 {
		fmt.Println("Mover disco de", origem, "para", destino)
		return
	}

	hanoi(n-1, origem, aux, destino)
	fmt.Println("Mover disco", n, "de", origem, "para", destino)
	hanoi(n-1, aux, destino, origem)
}

func main() {
	hanoi(3, "A", "C", "B")
}
