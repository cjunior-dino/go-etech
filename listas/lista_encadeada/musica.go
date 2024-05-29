package main

type Node struct {
	titulo  string
	artista string
	next    *Node
	prev    *Node
}

func (m music_List) Append(titulo string, artista string, next *Node) {
	node := &Node{titulo: titulo, artista: artista, next: next}
	if m.head == nil {
		m.head = node
		m.tail = node
	}

}

type music_List struct {
	head *Node
	tail *Node
}
