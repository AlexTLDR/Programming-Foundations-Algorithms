package main

import "fmt"

func main() {
	items := []string{"apple", "pear", "orange", "banana", "apple", "orange", "apple", "pear", "banana",
		"orange", "apple", "kiwi", "pear", "apple", "orange"}

	printUniqueValue(items)
}

func printUniqueValue(items []string) {
	dict := make(map[string]int)
	for _, num := range items {
		dict[num] = dict[num] + 1
	}
	fmt.Println(dict)
}
