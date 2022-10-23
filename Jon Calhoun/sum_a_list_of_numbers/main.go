package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(Sum(numbers))
	fmt.Println(SumRecursive(numbers))
}

func Sum(numbers []int) int {
	var sum int

	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumRecursive(numbers []int) int {

	if len(numbers) == 0 {
		return 0
	}

	return numbers[0] + SumRecursive(numbers[1:])
}
