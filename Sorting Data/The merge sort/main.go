package main

import "fmt"

func main() {
	items := []int{6, 20, 8, 19, 56, 23, 87, 41, 49, 53}
	fmt.Println("unsorted:", items)
	fmt.Println("sorted:", mergeSort(items))
}

func mergeSort(dataset []int) []int {
	if len(dataset) < 2 {
		return dataset
	}
	leftarr := mergeSort(dataset[:len(dataset)/2])
	rightarr := mergeSort(dataset[len(dataset)/2:])
	return merge(leftarr, rightarr)

}

func merge(leftarr, rightarr []int) []int {
	merged := []int{}
	i := 0
	j := 0

	for i < len(leftarr) && j < len(rightarr) {
		if leftarr[i] < rightarr[j] {
			merged = append(merged, leftarr[i])
			i++
		} else {
			merged = append(merged, rightarr[j])
			j++
		}
	}
	for ; i < len(leftarr); i++ {
		merged = append(merged, leftarr[i])
	}
	for ; j < len(rightarr); j++ {
		merged = append(merged, rightarr[j])
	}
	return merged
}
