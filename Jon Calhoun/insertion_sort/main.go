package main

import "fmt"

func main() {
	list := []int{54, 32, 786, 90, 12, 37, 89}
	InsertionSortInt(list)
	fmt.Println(list)
}

func InsertionSortInt(list []int) {
	var sorted []int
	for _, item := range list {
		sorted = insert(sorted, item)
	}
	for i, val := range sorted {
		list[i] = val
	}
}

func insert(sorted []int, item int) []int {
	for i, sortedItem := range sorted {
		if item < sortedItem {
			// sorted[:i] + item + sorted[i:]

			/*
				from https://github.com/golang/go/wiki/slicetricks using the insert

				a = append(a[:i], append([]T{x}, a[i:]...)...)

				endList := append([]int{item}, sorted[i:]...)
				sorted := append(sorted[:i], endList...)
			*/

			// sorted = append(sorted[:i], append([]int{item}, sorted[i:]...)...)
			return append(sorted[:i], append([]int{item}, sorted[i:]...)...)
		}
	}
	return append(sorted, item)
}
