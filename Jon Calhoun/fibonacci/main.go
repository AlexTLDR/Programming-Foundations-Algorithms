package main

import "fmt"

func main() {
	fmt.Println(Fibonacci(14))
}

func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	/* recursive solution - because of recursion, this solution is very slow

	return Fibonacci(n-1) + Fibonacci(n-2)

	*/

	//for loop solution

	nMinus2 := 0
	nMinus1 := 1
	var current int
	for i := 2; i <= n; i++ {
		current = nMinus2 + nMinus1
		nMinus2 = nMinus1
		nMinus1 = current
	}
	return current
}
