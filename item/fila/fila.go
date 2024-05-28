package main

import "fmt"

type Fila struct {
	nome []string
}

func (ln *Fila) isEmpty() bool {
	return len(ln.nome) == 0
}

func (q *Fila) enqueue(nome string) {
	q.nome = append(q.nome, nome)
}

func (f *Fila) display() int {
	if f.isEmpty() == true {
		return 0
	} else {
		return len(f.nome)
	}
}

func main() {
	fila := &Fila{}
	var n string
	var x int
	for {
		fmt.Printf("INFORME AS OPÇOES\n")
		fmt.Printf("1_VERIFICAR A FILA\n2_ENTRAR NA FILA\n3_ATENDER\n")
		fmt.Scanln(&x)
		if x == 1 {
			if fila.isEmpty() {
				fmt.Printf("FILA VAZIA\n")
			} else {
				fmt.Printf("EXISTEM %d NA FILA\n", fila.display())
			}
		} else if x == 2 {
			fmt.Println("INFORME O SEU NOME\n")
			fmt.Scanln(&n)
			fila.enqueue(n)
		} else if x == 3 {
			fmt.Println("VAI SER IMPLEMENTADO\n")
		}
	}
}
