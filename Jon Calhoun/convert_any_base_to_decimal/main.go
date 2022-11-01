package main

import (
	"fmt"
)

func main() {
	value := "1110"
	base := 2

	fmt.Println(BaseToDec(value, base))
}

func BaseToDec(value string, base int) int {
	result := 0
	multiplier := 1
	var val int
	for i := len(value) - 1; i >= 0; i-- {
		fmt.Sscanf(string(value[i]), "%X", &val)
		result = result + multiplier*val
		multiplier = multiplier * base
	}
	return result
}
