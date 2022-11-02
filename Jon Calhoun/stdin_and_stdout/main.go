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

	fmt.Println("----------------------------------------------")
	fmt.Println("Please enter the number of imputs with space between them")
	var m int
	fmt.Scanf("%d", &m)

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scanf("%d %d", &a, &b)
		fmt.Println("The GCD is:", GCD(a, b))

	}
}

func GCD(a, b int) int {

	for b != 0 {
		a, b = b, a%b
	}
	return a
}
