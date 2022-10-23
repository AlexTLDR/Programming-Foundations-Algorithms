package main

import "fmt"

func main() {
	items := []string{"apple", "pear", "orange", "banana", "apple", "orange", "apple", "pear", "banana",
		"orange", "apple", "kiwi", "pear", "apple", "orange"}

	fmt.Println(removeDuplicate(items))
}

func removeDuplicate(items []string) []string {

	// create a hashtable
	filter := make(map[string]bool)
	list := []string{}

	//loop over each item
	for _, item := range items {
		if _, value := filter[item]; !value {
			filter[item] = true
			list = append(list, item)
		}
	}
	return list
}
