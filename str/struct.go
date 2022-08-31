package main

import "fmt"

func main() {
	str := "Hello, 월드"

	runes := []rune(str)

	fmt.Println(len(str))
	fmt.Println(len(runes))
}