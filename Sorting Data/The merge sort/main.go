package main

import "fmt"

func main() {
	items := []int{6, 20, 8, 19, 56, 23, 87, 41, 49, 53}
	fmt.Println("unsorted:", items)
	mergeSort(items)
	fmt.Println("sorted:", items)
}

func mergeSort(dataset []int) []int {
	if len(dataset) > 1 {
		mid := len(dataset) / 2
		leftarr := dataset[:mid]
		rightarr := dataset[mid:]

		// recursively break down the arrays
		mergeSort(leftarr)
		mergeSort(rightarr)
		return merge(leftarr, rightarr)

	
}

func merge(leftarr, rightarr []int) []int {
	// perform the merging

	i := 0 // index into the left array
	j := 0 // index into the right array
	k := 0 // index into the merged array

	// while both arrays have content

	for i < len(leftarr) && j < len(rightarr) {
		if leftarr[i] < rightarr[j] {
			dataset[k] = leftarr[i]
			i += 1
		} else {
			dataset[k] = rightarr[j]
			j += 1
			k += 1
		}
	}

	// if the left array still has values, add them

	for i < len(leftarr) {
		dataset[k] = leftarr[i]
		i += 1
		k += 1
	}

	// if the right array still has values, add them

	for j < len(rightarr) {
		dataset[k] = rightarr[j]
		j += 1
		k += 1
	}

}
