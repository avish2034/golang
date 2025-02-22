package main

import (
	"fmt"
	"maps"
)

func main() {

	// m := make(map[string]string)

	// m["JS"] = "Javascript"
	// m["PY"] = "Python"

	// fmt.Println(m["JS"])
	// fmt.Println(m["PY"])

	// delete(m, "JS")
	// fmt.Println(m["JS"] == "")

	// clear(m)
	// fmt.Println(m)

	m := map[string]int{"price": 89, "item": 3}
	m1 := map[string]int{"price": 89, "item": 3}
	fmt.Println(m["price"])

	key, value := m["price"]

	fmt.Println(key, value)

	if value {
		fmt.Printf("VALUE \n")
	} else {
		fmt.Printf("KEY")
	}
	fmt.Println(maps.Equal(m, m1)) // For Copare Two Maps
}
