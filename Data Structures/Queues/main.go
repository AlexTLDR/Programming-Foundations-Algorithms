package main

import "fmt"

type queue struct {
	values []int
}

func main() {
	q := createQueue()
	q.add(1)
	q.add(2)
	q.add(3)
	q.add(4)
	fmt.Println(*q)
}

//Creating the Queue based on the queue struct

func createQueue() *queue {
	return &queue{
		values: []int{},
	}
}

//Adding elements to the Queue

func (q *queue) add(n int) *queue {
	q.values = append(q.values, n)
	return q
}
