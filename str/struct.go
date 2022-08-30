package main

import "fmt"

type Data struct {
	Name string
	Age  int
}


func main() {
	p := &Data{}
	x := new(Data)

	fmt.Println(&p,&x)
}