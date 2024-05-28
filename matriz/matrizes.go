package main

import "fmt"

func main() {

	var matriz [3][3]int

	contador := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matriz[i][j] = contador
			contador++
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			fmt.Printf("[%d] ", matriz[i][j])
		}
		fmt.Printf("\n")

	}
}
