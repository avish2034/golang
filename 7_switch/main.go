package main

import (
	"fmt"
	"time"
)

func main() {
	//simple switch

	i := 7

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Default")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Yooha , It's a weekend")
	}

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Its an integer")
		case string:
			fmt.Println("Its an string")
		case bool:
			fmt.Println("Its an Boolean")
		default:
			fmt.Println("Other", t)
		}
	}
	whoAmI(2.5)

}
