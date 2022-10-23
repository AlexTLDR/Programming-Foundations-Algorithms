package main

import "fmt"

func main() {
	items := []int{6, 8, 19, 20, 23, 41, 49, 53, 56, 87}
	fmt.Println(binarySearch(23, items))
	fmt.Println(binarySearch(87, items))
	fmt.Println(binarySearch(250, items))
}

func binarySearch(item int, itemList []int) any {
	listsize := len(itemList) - 1

	//start at the 2 ends of the list
	lowerIdx := 0
	upperIdx := listsize

	for lowerIdx <= upperIdx {
		// calculate the middle point
		midPt := (lowerIdx + upperIdx) / 2

		// if item is found return the index

		if itemList[midPt] == item {
			return midPt
		}

		// otherwise get the next midpoint

		if item > itemList[midPt] {
			lowerIdx = midPt + 1
		} else {
			upperIdx = midPt - 1
		}

		// if lowerIdx > upperIdx {

		// }
	}
	return "None"
}
