package main

import (
	"fmt"
	"iter"
)

type Numbers []int

func (n *Numbers) PrintEventNumbers() iter.Seq[*Number]

type Number int

func (nn *Number) IsEven() bool {
	return *nn%2 == 0
}

func (nn *Number) Value() Number {
	return *nn
}

var numbers = Numbers{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func main() {
	for number := range numbers.PrintEventNumbers() {
		if number.IsEven() {
			fmt.Println(number.Value())
		}
	}
}
