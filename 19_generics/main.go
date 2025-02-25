package main

import "fmt"

//you can add multiple generic types
func printSlice[T comparable, V string](items []T) { // [T any ] , [T interface{}]      or use like       [T string | int]
	for _, item := range items {
		fmt.Println(item)
	}
}

// create Stack LIFO

type stack[T any] struct { // Give Generic Type in Struct
	element []T
}

func main() {
	printSlice([]int{1, 2, 3})

	myStack := stack[string]{
		element: []string{"golang"},
	}
	fmt.Println(myStack)
}
