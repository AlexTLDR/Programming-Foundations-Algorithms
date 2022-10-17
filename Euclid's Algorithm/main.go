package main

import "fmt"

func main() {
	fmt.Println(gcd(60, 96))
	fmt.Println(gcd(20, 8))
}

func gcd(a, b int) int {
	for b != 0 {
		t := a
		a = b
		b = t % b

	}
	return a
}
