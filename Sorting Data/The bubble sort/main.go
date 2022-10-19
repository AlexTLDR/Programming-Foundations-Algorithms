package main

import "fmt"

func main() {
	list1 := []int{6, 20, 8, 19, 56, 23, 87, 41, 49, 53}
	bubbleSort(list1)
	fmt.Println("Result:", list1)
}

func bubbleSort(dataset []int) {

	for i := range dataset[1:] {
		if dataset[i] > dataset[i+1] {
			temp := dataset[i]
			dataset[i] = dataset[i+1]
			dataset[i+1] = temp
		}
		fmt.Println("Current state:", dataset)
	}

}
