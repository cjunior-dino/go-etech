// Esta aplicação suporta as operações básicas de uma fila:
//Enqueue (adicionar um elemento ao final da fila),
//Dequeue (remover um elemento do início da fila),
//Front (verificar o primeiro elemento da fila)
//IsEmpty (verificar se a fila está vazia)
// Display mostrar todos os valores na fila.

package main

import (
	"container/list"
	"fmt"
	"net/http"
)

type Queue struct {
	v *list.List
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func (q *Queue) Enqueue(val interface{}) {
	q.v.PushBack(val)
}

func (q *Queue) Dequeue() interface{} {
	ele := q.v.Front()
	if ele != nil {
		q.v.Remove(ele)
		return ele.Value
	}
	return nil
}

func (q *Queue) Front() interface{} {
	ele := q.v.Front()
	if ele != nil {
		return ele.Value
	}
	return nil
}

func (q *Queue) IsEmpty() bool {
	return q.v.Len() == 0
}

func (q *Queue) Display() string {
	var result string
	for e := q.v.Front(); e != nil; e = e.Next() {
		result += fmt.Sprint(e.Value) + "\n"
	}
	return result
}

func main() {
	q := NewQueue()

	http.HandleFunc("/enqueue", func(w http.ResponseWriter, r *http.Request) {
		val := r.URL.Query().Get("val")
		q.Enqueue(val)
	})

	http.HandleFunc("/dequeue", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, q.Dequeue())
	})

	http.HandleFunc("/front", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, q.Front())
	})

	http.HandleFunc("/isempty", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, q.IsEmpty())
	})

	http.HandleFunc("/display", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, q.Display())
	})

	// Adicione esta linha para servir arquivos estáticos a partir da pasta Templates
	http.Handle("/", http.FileServer(http.Dir("./Templates")))

	http.ListenAndServe(":4200", nil)
}
