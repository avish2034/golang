package main

import "fmt"

func main() {
	const name string = "avish"
	fmt.Println(name)
	// name = "new"

	// fmt.Println(name)

	//multiple constant
	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(host, ":", port)

}
