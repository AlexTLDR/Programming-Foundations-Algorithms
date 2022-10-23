package main

import "fmt"

func main() {
	list1 := []int{6, 20, 8, 19, 56, 23, 87, 41, 49, 53}
	bubbleSort(list1)
	fmt.Println("Result:", list1)
}

func bubbleSort(dataset []int) {
	for i := len(dataset); i > 0; i-- {
		for j := 1; j < i; j++ {
			if dataset[j-1] > dataset[j] {
				temp := dataset[j]
				dataset[j] = dataset[j-1]
				dataset[j-1] = temp
				fmt.Println("Current state:", dataset)
			}
		}
	}

}
