package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	lenght int
	head   *node
	tail   *node
}

func main() {

	list := linkedList{}
	node1 := &node{data: 38}
	node2 := &node{data: 49}
	//node3 := &node{data: 13}
	//node4 := &node{data: 15}

	list.PushBack(node1)
	list.PushBack(node2)
	l := linkedList(list)
	fmt.Println("l = ", l)
	fmt.Println(node1)

}

// Lenght of the Linked List

func (l linkedList) Len() int {
	return l.lenght
}

// Linked List PushBack - The PushBack Method takes a node as input and links it to the linked list.

func (l *linkedList) PushBack(n *node) {
	if l.head == nil {
		l.head = n
		l.tail = n
		l.lenght++
	} else {
		l.tail.next = n
		l.tail = n
		l.lenght++
	}
}
