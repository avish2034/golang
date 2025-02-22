package main

import (
	"fmt"
)

func main() {
	//unintialized slice is nil
	// var nums []int

	// fmt.Println(nums == nil) // used nil behalf of null

	var nums = make([]int, 0, 5)
	// nums := []int{1, 2, 3, 4, 5}
	nums = append(nums, 4, 2, 3, 4, 5, 2)
	nums2 := make([]int, len(nums))

	//COPY Operator

	copy(nums2, nums) //(destintion,source)

	//slice operator

	// fmt.Println(nums[0:1]) //[start:end] output will be end-1

	// silce Equal
	// fmt.Println(slices.Compare(nums, nums2))

	fmt.Println(nums, cap(nums))
}
