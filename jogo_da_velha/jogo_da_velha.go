package main

import "fmt"

var jogo [][] string{
	{" "," "," "},
	{" "," "," "},
	{" "," "," "},
	

}
var cont = 'X'

func marca(jogada [][]string) {
	x := 0
	y := 0
	fmt.Println("Informe a linha")
	fmt.Scanln(&x)
	fmt.Println("Informe a coluna")
	fmt.Scanln(&y)

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {

			if i == x && j == y {
				fmt.Scanln(&jogada[i][j])
			}

		}

	}
}

func main() {

	marca(jogo[][])

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(" ", jogo[i][j])
		}
		fmt.Printf("\n")
	}

}
