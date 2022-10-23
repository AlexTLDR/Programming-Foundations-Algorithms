package main

import "fmt"

func main() {
	numbers := []int{}
	fmt.Println(Sum(numbers))
}

func Sum(numbers []int) int {
	var sum int

	for _, v := range numbers {
		sum += v
	}
	return sum
}
