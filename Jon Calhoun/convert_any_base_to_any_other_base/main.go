package main

import "fmt"

func main() {
	fmt.Println(BaseToBase("E", 16, 2))
}

func BaseToBase(value string, base, newBase int) string {

	return DecToBase(BaseToDec(value, base), newBase)
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

func DecToBase(dec, base int) string {
	var res string
	for dec > 0 {
		rem := dec % base
		res = fmt.Sprintf("%X%s", rem, res)
		dec = dec / base
	}
	return res

}
