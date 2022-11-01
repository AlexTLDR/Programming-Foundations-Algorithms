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
	//i, _ := strconv.Atoi(value)
	result := 0
	for i := 0; i < len(value); i++ {
		digit := base ^ len(value) - 1
		result += digit
	}
	return result
}
