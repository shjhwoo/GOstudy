package B

import (
	"fmt"
)

type virA interface {
	SayFromA()
}

type B struct {
	A virA
}

func NewB(a virA) *B {
	return &B{
		A: a,
	}
}

func (b *B) SayFromB() {
	fmt.Println("B say Hello")
}

func (b *B) SayFromA() {
	// tmp_a := A.NewA()
	// tmp_a.SayFromA()
	b.A.SayFromA()
}

//https://incredible-larva.tistory.com/entry/err-import-cycle-not-allowed
