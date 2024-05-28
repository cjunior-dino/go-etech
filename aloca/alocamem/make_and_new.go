package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	s := make([]int, 1, 10)
	fmt.Println(s)
	elapsed := time.Since(start)

	fmt.Printf("Tempo gasto com a capacidade inicial definida: %s\n", elapsed)
	p := new(int)
	fmt.Println(*p)

	elapsed = time.Since(start)

	fmt.Printf("Tempo gasto com a capacidade inicial definida: %s\n", elapsed)
}
