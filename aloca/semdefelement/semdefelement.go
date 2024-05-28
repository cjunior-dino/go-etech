package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	slice := make([]int, 0)

	for i := 0; i < 1000000; i++ {
		slice = append(slice, i)
	}

	elapsed := time.Since(start)

	fmt.Printf("Tempo gasto sem definir a capacidade inicial: %s\n", elapsed)
}
