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
	s.Pop()
	fmt.Printf("stack after Pop method: %v\n", *s)
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

//Push adds an element to the stack

func (s *stack) Push(n int) *stack {
	s.count = append(s.count, n)
	return s
}

//Pop removes the last element from the stack and returns the rest of the elements

func (s *stack) Pop() (int, error) {
	if len(s.count) == 0 {
		return 0, fmt.Errorf("empty stack")
	}
	fmt.Println(s.count[len(s.count)-1]) //print out the removed element
	element := s.count[len(s.count)-1]
	s.count = s.count[:len(s.count)-1]
	return element, nil
}
