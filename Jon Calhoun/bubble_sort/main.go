package main

import "fmt"

func main() {

	list := []int{78, 34, 98, 10, 17}
	BubbleSortInt(list)
	fmt.Println(list)

}

func BubbleSortInt(list []int) {
	for i := 0; i < len(list)-1; i++ {
		swapped := false
		for j := 1; j < i; j++ {
			if list[j-1] > list[j] {
				//swap - using the swapped bool to stop the loop if no swapping is necessary for optimization
				list[j-1], list[j] = list[j], list[j-1]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
