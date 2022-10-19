package main

import "fmt"

type node struct {
	value int
}

type stack struct {
	nodes []*node
	count int
}

func main() {
	s := newStack()
	s.Push(&node{1})
	s.Push(&node{2})
	s.Push(&node{3})
	s.Push(&node{4})
	fmt.Println(s)
}

//newStack returns a new stack

func newStack() *stack {
	return &stack{}
}

//Push adds a node to the stack

func (s *stack) Push(n *node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}
