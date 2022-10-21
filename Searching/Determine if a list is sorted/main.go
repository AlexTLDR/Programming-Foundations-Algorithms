package main

import "fmt"

func main() {
	items1 := []int{6, 8, 19, 20, 23, 41, 49, 53, 56, 87}
	items2 := []int{6, 20, 8, 19, 56, 23, 87, 41, 49, 53}

	fmt.Println(isSorted(items1))
	fmt.Println(isSorted(items2))

}

func isSorted(itemList []int) bool {

	//using the brute force method

	for i := range itemList {
		if itemList[i] > itemList[i+1] {
			return false
		}
	}
	return true
}
