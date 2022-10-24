package main

import "fmt"

func main() {
	n := 100
	FizzBuzz(n)
}

func FizzBuzz(n int) {

	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			fmt.Printf("Fizz Buzz")
		case i%5 == 0:
			fmt.Printf("Buzz")
		case i%3 == 0:
			fmt.Printf("Fizz")
		default:
			fmt.Printf("%d", i)
		}
		if i != n {
			fmt.Printf(", ")
		}

		// if i%15 == 0 {
		// 	fmt.Printf("Fizz Buzz")
		// } else if i%5 == 0 {
		// 	fmt.Printf("Buzz")
		// } else if i%3 == 0 {
		// 	fmt.Printf("Fizz")
		// } else {
		// 	fmt.Printf("%d", i)
		// }
		// if i < n {
		// 	fmt.Printf(", ")
		// } else {
		// 	fmt.Printf("\n")
		// }
	}
	fmt.Printf("\n")
}
