package main

func main() {
	items := []int{20, 6, 8, 53, 56, 23, 87, 41, 49, 19}

}

func quickSort(dataset []int, first, last int) {
	if first < last {
		pivotIdx := partition(dataset, first, last)

		quickSort(dataset, first, pivotIdx-1)
		quickSort(dataset, pivotIdx+1, last)
	}
}

func partition(datavalues []int, first, last int) int {

}
