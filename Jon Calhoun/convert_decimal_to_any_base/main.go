package main

import "fmt"

func main() {
	dec := 10
	base := 2

	fmt.Println(DecToBase(dec, base))
}

func DecToBase(dec, base int) string {
	var res string
	for dec != 0 {
		res = res + string(dec/base)
		dec = dec / base
	}

	return res
}
