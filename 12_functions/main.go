package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func langauge() (string, string, string) {
	return "Hindi", "English", "Gujarati"
}
func main() {
	fmt.Println(add(5, 6))
	fmt.Println(langauge())
	lang1, lang2, lang3 := langauge()
	fmt.Println(lang1, lang2, lang3)
}
