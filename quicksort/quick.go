package main

import (
	"fmt"
)

func quickSort(arr []int) []int {

	if len(arr) <= 1 {

		return arr
	}

	pivot := arr[len(arr)/2]
	var less, greater []int

	for _, v := range arr {
		if v < pivot {
			less = append(less, v)
		} else if v > pivot {
			greater = append(greater, v)
		}
	}

	less = quickSort(less)
	greater = quickSort(greater)

	return append(append(less, pivot), greater...)
}

func main() {
	arr := []int{5, 3, 8, 2, 7, 1, 9, 4, 6}
	fmt.Println(quickSort(arr))
}
