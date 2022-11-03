package main

import "fmt"

func main() {

	list := []int{78, 34, 98, 10, 17}
	BubbleSortInt(list)
	fmt.Println(list)

}

func BubbleSortInt(list []int) {
	for i := 0; i < len(list)-1; i++ {
		for j := 1; j < i; j++ {
			if list[j-1] > list[j] {
				//swap
				list[j-1], list[j] = list[j], list[j-1]
			}
		}
	}
}
