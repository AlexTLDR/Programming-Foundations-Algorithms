package main

import "fmt"

func main() {
	var n int
	fmt.Println("How many names")
	fmt.Scanf("%d", &n)
	fmt.Println("The required number of names is", n)
	var names []string
	for i := 0; i < n; i++ {
		fmt.Printf("Enter name %d:\n", i+1)
		var name string
		fmt.Scanf("%s\n", &name)
		names = append(names, name)
	}
	fmt.Println(names)
}
