package main

import "fmt"

func main() {
	countdown(5)
}

func countdown(x int) {
	if x == 0 {
		fmt.Println("Done!")
	} else {
		fmt.Println(x, "...")
		countdown(x - 1)
	}
}
