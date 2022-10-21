package main

import "fmt"

func main() {
	items := []int{6, 20, 8, 19, 56, 23, 87, 41, 49, 53}
	fmt.Println(findItem(87, items))
	fmt.Println(findItem(250, items))

}

func findItem(item int, itemList []int) any {
	for i := range itemList {
		if item == itemList[i] {
			return i
		}

	}
	return "none"

}
