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
		fmt.Println("foo")
		/* the foo will get printed x times as it is part of the countdown function
		5 ...
		4 ...
		3 ...
		2 ...
		1 ...
		Done!
		foo
		foo
		foo
		foo
		foo
		*/
	}
}
