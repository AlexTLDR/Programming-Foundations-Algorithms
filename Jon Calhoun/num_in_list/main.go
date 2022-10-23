package main

import "fmt"

func main() {
	num := 75
	list := []int{13, 42, 37, 24, 75}

	fmt.Println(NumInList(list, num))

}

func NumInList(list []int, num int) bool {

	for _, v := range list {
		if num == v {
			return true
		}
	}
	return false
}
