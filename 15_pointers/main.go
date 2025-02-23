package main

import "fmt"

func changeNum(num *int) {
	*num = 5 // dereference the value
	fmt.Println("In ChangeNum", *num)
}

func main() {
	num := 1

	changeNum(&num) // give the value reference

	fmt.Println("After ChangeNum in main", num)
}
