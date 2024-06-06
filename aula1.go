package main

import "fmt"

func main() {
	var x int
	fmt.Println("Hello, World!")
	fmt.Scanln(&x)
	if x%3 == 0 {
		fmt.Println("è muitiplo")
	} else {
		fmt.Println("Não é multiplo")
	}
}
