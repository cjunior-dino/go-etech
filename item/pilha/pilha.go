package main

import "fmt"

type Pilha struct {
	items []string
}

func (p *Pilha) push(item string) {
	p.items = append(p.items, item)
}

func (p *Pilha) pop() string {
	if len(p.items) == 0 {
		return ""
	}
	item := p.items[len(p.items)-1]
	p.items = p.items[:len(p.items)-1]
	return item

}

func (p *Pilha) peek() string {
	if len(p.items) == 0 {
		return ""
	}
	return p.items[len(p.items)-1]
}

func (p *Pilha) isEmpty() bool {
	return len(p.items) == 0
}

func main() {

	pilha := &Pilha{}

	pilha.push("teste")
	pilha.push("oi")
	pilha.push("aula")

	fmt.Println(pilha.items)

	fmt.Println(pilha.pop())

	fmt.Println(pilha.items)

}
