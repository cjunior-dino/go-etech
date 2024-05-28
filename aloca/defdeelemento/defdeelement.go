package main

import (
	"fmt"

	"time"
)

func main() {
	start := time.Now()

	slice := make([]int, 0, 1000000)

	for i := 0; i < 1000000; i++ {
		slice = append(slice, 1)
	}

	elapsed := time.Since(start)

	fmt.Printf("Tempo gasto com a capacidade inicial definida: %s\n", elapsed)
}
