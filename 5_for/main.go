package main

//for -> only construct in go for looping
func main() {
	i := 1
	for i <= 3 { // it is working like while loop
		println(i)
		i++
	}

	//infine loop
	// for {
	// 	println("i")
	// }

	//classic for loop
	// for i := 0; i < 3; i++ {
	// 	if i < 1 {
	// 		continue
	// 	}
	// 	println(i)
	// 	// break
	// }

	//range

	for i := range 3 { // range work like python it will go 0 to 2
		println(i)
	}
}
