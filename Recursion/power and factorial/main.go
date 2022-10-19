package main

import "fmt"

func main() {

	fmt.Println("5 to the power of 3 is", power(5, 3))
	fmt.Println("1 to the power of 5 is", power(1, 5))
	fmt.Println("4! is", factorial(4))
	fmt.Println("0! is", factorial(0))
}

func power(num, pwr int) int {
	if pwr == 0 {
		return 1
	} else {
		return num * power(num, pwr-1)
	}
}

func factorial(num int) int {
	if num == 0 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}
