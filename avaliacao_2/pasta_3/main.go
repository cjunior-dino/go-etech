package main

import (
	"fmt"
)

type Queue struct {
	data []int
}
type Queuepriority struct {
	data []int
}

func (q *Queue) Enqueue(val int) { // adiciona elemento na fila
	q.data = append(q.data, val)
}

func (q *Queue) Dequeue() int { // remove o primeiro elemento da fila
	if len(q.data) == 0 {
		return -1
	}
	val := q.data[0]
	q.data = q.data[1:]
	return val
}

func (q *Queue) Front() int { // exibe o primeiro elemento da fila
	if len(q.data) == 0 {
		return -1
	}
	return q.data[0]
}

func (q *Queue) IsEmpty() bool { // verifica se a fila esta vazia
	return len(q.data) == 0
}

func (q *Queuepriority) Enqueuepriority(val int) { // adiciona elemento na fila priority
	q.data = append(q.data, val)
}

func (q *Queuepriority) Dequeuepriority() int { // remove o primeiro elemento da fila priority
	if len(q.data) == 0 {
		return -1
	}
	val := q.data[0]
	q.data = q.data[1:]
	return val
}

func (q *Queuepriority) Frontpriority() int { // exibe o primeiro elemento da fila priority
	if len(q.data) == 0 {
		return -1
	}
	return q.data[0]
}

func (q *Queuepriority) IsEmptypriority() bool { // verifica se a fila esta vazia priority
	return len(q.data) == 0
}

func main() {
	fila := 1
	atende := 1
	queue := Queue{}
	priority := Queuepriority{}
	var x int

	for {
		fmt.Println("BEM-VINDO AO HOSPILTAL SAUDE LEGAL")
		fmt.Println(" ")
		fmt.Println("Escolha uma das opções")
		fmt.Println("1- Entrar na fila:")
		fmt.Println("2- Atender fila:")
		fmt.Println("3- Verificar o proximo:")
		fmt.Println("4- Sair do sistema:")
		fmt.Scanln(&x)

		if x == 1 {
			var y int
			fmt.Println("Voce é :")
			fmt.Println("1- Regular")
			fmt.Println("2- Prioridade")
			fmt.Scanln(&y)
			if y == 1 {
				queue.Enqueue(fila)
				fmt.Printf("O N° %d foi adicionado a fila normal\n", fila)
				fila++
			} else if y == 2 {
				priority.Enqueuepriority(fila)
				fmt.Printf("O N° %d foi adicionado a fila de prioridade\n", fila)
				fila++
			}
		} else if x == 2 {
			if atende%3 != 0 {
				if queue.IsEmpty() {
					fmt.Println(" ")
					atende++
				} else {
					fmt.Printf("Atendendo o N° %d da fila normal\n", queue.Dequeue())
					fmt.Println(" ")
					atende++
				}
			} else {
				if priority.IsEmptypriority() {
					fmt.Println(" ")
					atende++
				} else {
					fmt.Printf("Atendendo o N° %d da fila prioridade\n", priority.Dequeuepriority())
					fmt.Println(" ")
					atende++

				}

			}
		} else if x == 3 {
			if atende%3 != 0 {
				fmt.Printf("O proximo da fila é o N° %d da regular\n", queue.Front())
				fmt.Printf(" ")
			} else {
				fmt.Printf("O proximo da fila é o N° %d da prioridade\n", priority.Frontpriority())
				fmt.Println(" ")
			}
		} else if x == 4 {
			fmt.Println(" ")
			fmt.Println("Sistema encerrado !!!")
			break
		} else {
			fmt.Println(" ")
			fmt.Println("Opção invalida")
			fmt.Println(" ")
		}
	}

}
