package A

import (
	"cycle/B"
	"fmt"
)

type A struct {
	Age  int64
	Name string
}

func NewA() *A {
	return &A{}
}

func (a *A) SayFromA() {
	fmt.Println("A say Hello")
}

func (a *A) SayFromB() {
	tmp_b := B.NewB(a)
	tmp_b.SayFromB()
}
