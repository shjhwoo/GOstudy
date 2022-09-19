package main

import "fmt"

func main() {
	var a [3]int = [3]int{1,2,4}
	//var p *int
	//p = &a[0]
	fmt.Println(&a[0],&a[1],&a[2])
}