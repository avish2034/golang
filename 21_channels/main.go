package main

import "fmt"

// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("Nums is", num)
// 	}
// }

// func sum(result chan int, num1 int, num2 int) {
// 	sum := num1 + num2
// 	result <- sum
// }

//goroutine synchronizer
// func task(done chan bool) {
// 	defer func() { done <- true }()
// 	fmt.Println("Process................")
// }

func main() {

	emailChan := make(chan string, 100)

	emailChan <- "hell@redbull.com"
	emailChan <- "hell@redbull.com"
	emailChan <- "hell@redbull.com"
	emailChan <- "hell@redbull.com"

	fmt.Println((<-emailChan))
	fmt.Println((<-emailChan))
	fmt.Println((<-emailChan))
	fmt.Println((<-emailChan))

	// result := make(chan int) // unbuffer channel
	// go sum(result, 4, 5)
	// res := <-result // it is a bloacking channel
	// fmt.Println(res)

	// numChan := make(chan int)
	// go processNum(numChan)

	// for { // this is use like a queue.
	// 	numChan <- rand.Intn(100)
	// }

	// numChan <- 5

	// messageChan := make(chan string)
	// messageChan <- "ping"

	// message := <-messageChan

	// fmt.Println(message)

	// done := make(chan bool)
	// go task(done)

	// <-done //block

}
