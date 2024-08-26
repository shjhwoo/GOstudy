package main

import (
	"fmt"
	"iter"
	/*
		왜 iter를 써야 하나?
	*/)

type Int struct {
	Val     int
	Visible bool
}

type Ints []Int

func (ii Ints) All(yield func(Int) bool) {
	for _, s := range ii {
		if !yield(s) {
			return
		}
	}
}

func Filter[V any](seq iter.Seq[V], f func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if f(v) && !yield(v) {
				return
			}
		}
	}
}

func Map[V any](seq iter.Seq[V], f func(V) V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			v = f(v)
			if !yield(v) {
				return
			}
		}
	}
}

type Person struct {
	Age   int
	Phone string
	Email string
}

type People map[string]Person

func (pi People) All(yield func(name string, person Person) bool) {
	for name, pInfo := range pi {
		if !yield(name, pInfo) {
			return
		}
	}
}

func (pi People) PrintName() iter.Seq[string] {
	return func(yield func(name string) bool) {
		for name := range pi {
			if !yield(name) {
				continue
			}
		}
	}
}

// FilterAgeUnder23 ===  iterator 이다, 즉 콜백함수인 yield에게 연속적으로 원소들을 전달하는 함수임
func (pi People) FilterAgeUnder23() iter.Seq[string] {
	return func(yield func(name string) bool) {
		for name, pInfo := range pi {
			if pInfo.Age >= 23 || !yield(name) { //yield는 필요 없어 보여도, 안 쓰면은 아얘 값 자체를 못 받음..;
				continue //return 하면은 23 이상인 순간 멈춰서 그 뒤의 22는 못 받음
			}
		}
	}
}

func MapFilter[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if f(k, v) && !yield(k, v) {
				return
			}
		}
	}
}

func MapMap[K, V any](seq iter.Seq2[K, V], f func(K, V) (K, V)) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			k, v = f(k, v)
			if !yield(k, v) {
				return
			}
		}
	}
}

//iter: 값 조회가 목적이다.
//mutation을 해야 한다면!

type Tree []Leaf

type Leaf struct {
	Val int
}

func (t *Tree) Positions() iter.Seq[int] {
	return func(yield func(el int) bool) {
		for pos := range *t {
			if !yield(pos) {
				return
			}
		}
	}
}

func (t *Tree) All() iter.Seq[*Leaf] {
	return func(yield func(el *Leaf) bool) {
		/*
			n Go, when you range over a slice, the loop variable (e.g., l in for _, l := range *t)
			is a copy of the element at the current index in the slice.
			Since l is a new variable on each iteration,
			taking the address of l only gives you the address of the temporary copy,
			not the actual element in the slice.
		*/
		// for _, l := range *t { //이렇게 하면, 항상 원본의 복사본만 받아온다. 원본 변경불가
		// 	if !yield(&l) {
		// 		return
		// 	}
		// }
		for i := range *t {
			if !yield(&(*t)[i]) { // Yield the pointer to the actual slice element
				return
			}
		}
	}
}

func (p *Leaf) Value() int {
	return p.Val
}

func (p *Leaf) Delete() {
	p.Val = 0
}

func (p *Leaf) Set(v int) {
	p.Val = v
}

func main() {
	var t = Tree{
		{
			Val: 12,
		},
		{
			Val: 192,
		},
		{
			Val: 120,
		},
		{
			Val: 8,
		},
		{
			Val: 14,
		},
	}

	// for p := range t.Positions() {
	// 	fmt.Println(p)
	// }

	for p := range t.All() {
		p.Set(2342648)
	}

	for p := range t.All() { //변하지 않고 그대로여. 왜?
		fmt.Println(p.Val)
	}

	//ints := Ints{{1, true}, {2, false}, {3, true}}

	//기본
	// for n := range ints.All {
	// 	fmt.Println(n)
	// }

	//짝수만 가져오기(직접 필터 함수 넣으면 리턴값이 없어서 range over func불가함. range over 가능한 함수 자체를 리턴하도록 할것)
	// for n := range Filter(ints.All, getEvenNumber) {
	// 	fmt.Println(n)
	// }

	// //모든 원소에 3을 더하기
	// for n := range Map(ints.All, add3) {
	// 	fmt.Println(n)
	// }

	//맵을 가지고 해보기
	// people := People{
	// 	"jane": {
	// 		Age:   22,
	// 		Phone: "123-456",
	// 	},
	// 	"park": {
	// 		Age:   23,
	// 		Phone: "923-134",
	// 	},
	// 	"bob": {
	// 		Age:   22,
	// 		Phone: "024-321",
	// 	},
	// }

	// for p, pInfo := range people.All {
	// 	fmt.Println(p, pInfo)
	// }
	// /////
	// // for elem := range people.PrintName() {
	// // 	fmt.Println(elem)
	// // }

	// for elem := range people.FilterAgeUnder23() {
	// 	fmt.Println(elem)
	// }

	// for p, pInfo := range MapFilter(people.All, getAgeUnder23) {
	// 	fmt.Println(p, pInfo)
	// }

	// for p, pInfo := range MapMap(people.All, addOneToAge) {
	// 	fmt.Println(p, pInfo)
	// }

}

func getAgeUnder23(name string, pInfo Person) bool {
	return pInfo.Age < 23
}

func addOneToAge(name string, pInfo Person) (string, Person) {
	pInfo.Age += 1
	return name, pInfo
}

func getEvenNumber(num Int) bool {
	return num.Val%2 == 0
}

func add3(num Int) Int {
	num.Val += 3

	return num
}
