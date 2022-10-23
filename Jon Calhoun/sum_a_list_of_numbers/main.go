package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(numbers))
}

func Sum(numbers []int) int {
	var sum int
	if len(numbers) == 0 {
		return 0
	}

	for _, v := range numbers {
		sum = sum + v
	}
	return sum
}
