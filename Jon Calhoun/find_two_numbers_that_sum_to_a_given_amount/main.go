package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	sum := 7

	fmt.Println(FindTwoThatSum(numbers, sum))
}

func FindTwoThatSum(numbers []int, sum int) (int, int) {
	for i, v := range numbers {
		for j, v2 := range numbers {
			if i == j {
				continue
			}
			if v+v2 == sum {
				return i, j
			}
		}
	}
	return -1, -1
}
