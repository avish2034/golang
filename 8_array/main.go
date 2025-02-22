package main

import "fmt"

// array is number sequenced of specific length
func main() {
	// golang add 0 default in array
	// var nums [4]int

	// nums[0] = 14

	// var bools [3]string

	// fmt.Println(len(nums), nums[0], bools)

	// to declare in single line
	// nums := [3]int{1, 2, 3}

	//2D Array
	nums := [2][2]int{{1, 2}, {3, 4}}

	fmt.Println(nums)

	// - fixed size , that is predictable
	// - Memory Optimization
	// - Constant time access
}
