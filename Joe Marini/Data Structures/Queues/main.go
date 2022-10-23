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
	q.remove()
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

//Check if the Queue is empty

func (q *queue) isEmpty() bool {
	return len(q.values) == 0
}

//Remove the first element from the queue -> dequeue

func (q *queue) remove() (int, error) {
	if len(q.values) == 0 {
		return 0, fmt.Errorf("empty queue")
	}
	element := q.values[0]
	fmt.Println(element)
	q.values = q.values[1:]
	return element, nil
}
