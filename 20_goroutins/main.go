package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	fmt.Println("Doing Task ", id)
	defer w.Done()
}

func main() {
	var waitGroup sync.WaitGroup
	for i := 0; i <= 10; i++ {
		waitGroup.Add(1)
		go task(i, &waitGroup)
	}
	waitGroup.Wait()
	// time.Sleep(time.Second * 2) //add sleep mode
}
