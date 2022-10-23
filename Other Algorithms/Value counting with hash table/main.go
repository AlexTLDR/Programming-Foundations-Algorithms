package main

import "fmt"

func main() {
	items := []string{"apple", "pear", "orange", "banana", "apple", "orange", "apple", "pear", "banana",
		"orange", "apple", "kiwi", "pear", "apple", "orange"}

	printUniqueValue(items)
}

func printUniqueValue(items []string) {

	// create a hashtable object to hold the items and counts

	dict := make(map[string]int)

	//iterate over each item and increment the count for each one

	for _, num := range items {
		dict[num] = dict[num] + 1
	}
	fmt.Println(dict)
}
