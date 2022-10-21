package main

import "fmt"

func main() {
	items := []int{6, 20, 8, 19, 56, 23, 87, 41, 49, 53}
	fmt.Println(findItem(87, items))
	fmt.Println(findItem(250, items))

}

func findItem(item int, itemList []int) any {
	for _, v := range itemList {
		if item == v {
			return v
		}

	}
	return "None"

}
