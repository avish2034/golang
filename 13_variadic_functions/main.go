package main

import "fmt"

func sum(nums ...int) int { // it will accept the unlimited values
	total := 0
	for _, num := range nums { // the params will access as a slices
		total += num
	}
	return total
}

func main() {
	fmt.Println(sum(2, 3, 4, 5, 6, 7, 78, 8, 8, 6, 6, 54, 44))
}
