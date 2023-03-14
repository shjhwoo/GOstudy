package main

import "fmt"

func main() {
	var array [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var slice1 []int = array[1:5]
	var slice2 []int = slice1[1:8:9] //기존 슬라이스가 아닌 기존 슬라이스가 가리키는 원본 배열을 가리키므로 원래 슬라이스의 범위를 넘어 슬라이스할수있다.
	var slice3 []int = make([]int, 5)
	var slice4 []int = make([]int, 0)
	var slice5 []int = []int{1, 2, 3, 4, 5}
	var slice6 []int

	fmt.Println(cap(array))

	fmt.Println(slice1, cap(slice1))
	fmt.Println(slice2, cap(slice2))
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)
	fmt.Println(slice6)

}
