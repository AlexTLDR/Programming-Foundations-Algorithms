package main

import "fmt"

func main() {
	items := []int{6, 20, 8, 19, 56, 23, 87, 41, 49, 53}
	fmt.Println("unsorted:", items)
	mergeSort(items)
	fmt.Println("sorted:", items)
}

func mergeSort(dataset []int) {
	if len(dataset) > 1 {
		mid := len(dataset) / 2
		leftarr := dataset[:mid]
		rightarr := dataset[mid:]

		// recursively break down the arrays

		mergeSort(leftarr)
		mergeSort(rightarr)

		// perform the merging

		i := 0
		j := 0
		k := 0

		// while both arrays have content

		if i < len(leftarr) && j < len(rightarr) {
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
		if i < len(leftarr) {
			dataset[k] = leftarr[i]
			i += 1
			k += 1
		}

		// if the right array still has values, add them

		if j < len(rightarr) {
			dataset[k] = rightarr[i]
			i += 1
			k += 1
		}
	}

}
