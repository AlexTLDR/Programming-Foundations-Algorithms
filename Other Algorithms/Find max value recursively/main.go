package main

import "fmt"

func main() {
	items := []int{6, 20, 8, 19, 56, 23, 87, 41, 49, 53}
	fmt.Println(findMax(items))

}

func findMax(items []int) int {
	if len(items) == 1 {
		return items[0]
	} else {
		op1 := items[0]
		op2 := findMax(items[1:])

		if op1 > op2 {
			return op1
		}
		return op2
	}

}
