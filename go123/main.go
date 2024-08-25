package main

import (
	"fmt"
	"iter"
)

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
				return
			}
		}
	}
}

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

func main() {
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
	people := People{
		"jane": {
			Age:   22,
			Phone: "123-456",
		},
		"park": {
			Age:   23,
			Phone: "923-134",
		},
		"bob": {
			Age:   22,
			Phone: "024-321",
		},
	}

	for p, pInfo := range people.All {
		fmt.Println(p, pInfo)
	}
	/////
	// for elem := range people.PrintName() {
	// 	fmt.Println(elem)
	// }

	for elem := range people.FilterAgeUnder23() {
		fmt.Println(elem)
	}

	for p, pInfo := range MapFilter(people.All, getAgeUnder23) {
		fmt.Println(p, pInfo)
	}

	for p, pInfo := range MapMap(people.All, addOneToAge) {
		fmt.Println(p, pInfo)
	}
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
