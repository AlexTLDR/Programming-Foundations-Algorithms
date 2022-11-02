package main

import "fmt"

func main() {
	fmt.Println(GCD(15, 9))
}

func GCD(a, b int) int {

	for b != 0 {
		a, b = b, a%b
	}
	return a
}
