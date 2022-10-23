package main

import "fmt"

func main() {
	word := "cat"
	fmt.Println(Reverse(word))
	fmt.Println(BuildStrReversed(word))
}

func Reverse(word string) string {
	var reversed string
	for i := len(word) - 1; i >= 0; i-- {
		reversed = reversed + string(word[i])
	}

	return reversed
}

func BuildStrReversed(word string) string {
	var reversed string
	for i := range word {
		reversed = string(word[i]) + reversed
	}

	return reversed
}
