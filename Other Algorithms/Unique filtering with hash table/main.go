package main

import "fmt"

func main() {
	items := []string{"apple", "pear", "orange", "banana", "apple", "orange", "apple", "pear", "banana",
		"orange", "apple", "kiwi", "pear", "apple", "orange"}

	//create a hashtable

	// filter := make(map[string]int)

	// for _, s := range items {
	// 	filter[s]++
	// }

	// fmt.Println(filter)

	fmt.Println(removeDuplicate(items))
}

func removeDuplicate(items []string) []string {
	filter := make(map[string]bool)
	list := []string{}
	for _, item := range items {
		if _, value := filter[item]; !value {
			filter[item] = true
			list = append(list, item)
		}
	}
	return list
}
