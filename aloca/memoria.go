package main

import "fmt"

func main() {
	slice := make([]int, 0, 10)

	for i := 0; i <= 10; i++ {
		slice = append(slice, i)
	}

	fmt.Println(slice)
}
