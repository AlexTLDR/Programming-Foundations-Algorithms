package main

import "fmt"

type stack struct {
	count []int
}

func main() {
	s := newStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Printf("%v\n", *s)
}

//newStack returns a new stack

func newStack() *stack {
	return &stack{
		count: []int{},
	}
}

func (s *stack) isEmpty() bool {
	return len(s.count) == 0
}

//Push adds a node to the stack

func (s *stack) Push(n int) *stack {
	s.count = append(s.count, n)
	return s
}
