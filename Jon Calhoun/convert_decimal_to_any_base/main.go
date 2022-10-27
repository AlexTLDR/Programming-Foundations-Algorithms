package main

import "fmt"

func main() {
	dec := 10
	base := 2

	fmt.Println(DecToBase(dec, base))
}

func DecToBase(dec, base int) string {
	var res string
	for dec > 0 {
		rem := dec % base
		res = fmt.Sprintf("%d%s", rem, res)
		dec = dec / base
	}
	return res

}
