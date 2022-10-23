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
	// for i := range word {
	// 	reversed = string(word[i]) + reversed
	// }

	// making the code test the rune test

	for _, r := range word {
		reversed = string(r) + reversed
	}

	return reversed
}
