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
	node1 := &node{data: 15}
	node2 := &node{data: 13}
	node3 := &node{data: 49}
	node4 := &node{data: 38}

	list.PushBack(node1)
	list.PushBack(node2)
	list.PushBack(node3)
	list.PushBack(node4)

	//l := linkedList(list)
	//fmt.Println("l = ", l)
	//fmt.Println(node1)

	list.Display()
	fmt.Println("Item count before delete:", list.Len())
	list.Delete(38)
	list.Display()
	fmt.Println("Item count after delete:", list.Len())

}

// Lenght of the Linked List

func (l linkedList) Len() int {
	return l.lenght
}

// Display the Linked List

func (l linkedList) Display() {
	for l.head != nil {
		fmt.Printf("Node: %v\n", l.head.data)
		l.head = l.head.next
	}
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

// Delete a Node from the Linked List

func (l *linkedList) Delete(key int) {
	if l.head.data == key {
		l.head = l.head.next
		l.lenght--
		return
	}

	var previous *node = nil
	current := l.head
	for current != nil && current.data != key {
		previous = current
		current = current.next
	}
	if current == nil {
		fmt.Println("Key not found")
		return
	}
	previous.next = current.next
	l.lenght--
	fmt.Println("Node deleted")

}
