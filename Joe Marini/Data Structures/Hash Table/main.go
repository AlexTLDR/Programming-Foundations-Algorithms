package main

import "fmt"

func main() {
	items1 := map[string]any{
		"key1": 1,
		"key2": 1,
		"key3": "three",
	}

	items2 := make(map[string]any)
	items2["key1"] = 1
	items2["key2"] = 2
	items2["key3"] = "three"

	fmt.Println(items1, items2)
	//fmt.Println(items1["key4"]) // in the course, in Python this retuns an error but in Go it returns nil
	if key, ok := items1["key4"]; !ok {
		fmt.Println("Key not found:", key)
	} else {
		fmt.Println(key)
	}
	items2["key2"] = "two"
	fmt.Println(items2)

	for k, v := range items2 {
		fmt.Println("key:", k, ",", "value:", v)
	}
}
