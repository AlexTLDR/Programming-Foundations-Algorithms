package main

import "fmt"

func main() {
	items := []int{20, 6, 8, 53, 56, 23, 87, 41, 49, 19}
	fmt.Println("before sort:", items)
	quickSort(items, 0, len(items)-1)
	fmt.Println("after sort:", items)

}

func quickSort(dataset []int, first, last int) {
	if first > last {
		return
	}
	pivotIdx := partition(dataset, first, last)

	quickSort(dataset, first, pivotIdx-1)
	quickSort(dataset, pivotIdx+1, last)

}

func partition(datavalues []int, first, last int) int {
	pivotvalue := datavalues[last]
	for i := first; i < last; i++ {
		if datavalues[i] < pivotvalue {
			datavalues[i], datavalues[first] = datavalues[first], datavalues[i]
			first++
		}
	}

	datavalues[first], datavalues[last] = datavalues[last], datavalues[first]
	return first
}
