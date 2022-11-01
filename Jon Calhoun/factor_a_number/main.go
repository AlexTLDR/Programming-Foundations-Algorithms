package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7}
	number := 30
	fmt.Println(Factor(primes, number))
}

func Factor(primes []int, number int) []int {
	var res []int
	for _, prime := range primes {
		for number%prime == 0 {
			res = append(res, prime)
			number = number / prime
		}
	}
	if number > 1 {
		res = append(res, number)
	}
	return res
}
