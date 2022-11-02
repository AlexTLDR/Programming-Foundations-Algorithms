package main

import "fmt"

func main() {
	fmt.Println(GCD(20, 10))
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
