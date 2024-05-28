package main

import (
	"fmt"
	"net/http"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func main() {
	product := Product{
		ID:    "1",
		Name:  "Product 1",
		Price: 100.00,
	}

	fmt.Println(product.Name)
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
