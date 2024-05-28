package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Stack []string

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(v string) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func main() {
	s := &Stack{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nDigite uma ação para adicionar à pilha ou 'desfazer' para desfazer a última ação:")
		scanner.Scan()
		text := scanner.Text()

		if strings.ToLower(text) == "desfazer" {
			action, success := s.Pop()

			if success {
				fmt.Println("Ação desfeita: ", action)
			} else {
				fmt.Println("Não há mais ações para desfazer.")
			}
		} else {
			s.Push(text)
			fmt.Println("Ação adicionada:", text)
		}
	}
}
