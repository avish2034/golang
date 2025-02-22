package main

import "fmt"

//iterating over data structure
func main() {

	//iterate on slices
	// nums := []int{6, 7, 8}

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }
	// var sum = 0
	// for _, num := range nums {
	// 	sum += num
	// }
	// fmt.Println(sum)

	//iterate on maps
	m := map[string]string{"name": "avish", "age": "20"}

	for key, value := range m {
		fmt.Println(key, value)
	}


	
}
