package main

import (
	"fmt"
	"sync"
	"time"
)

// func main() {
// 	var messages chan string = make(chan string) //채널명 채널타입 = make(메세지 타입), 메세지는 고루틴끼리 주고받는 정보!

// 	//채널에 데이터 넣기. 왼쪽에는 채널 인스턴스 명을 기재한다. 데이터를 넣을 때는 <- 연산자를 사용한다
// 	messages <- "this is a message"

// 	//채널에서 데이터를 빼기. 화살표 방향은, 데이터가 담기는 그릇을 향한다
// 	var msg string = <- messages
// }

func main() {
	var wg sync.WaitGroup
	ch := make(chan int) //나는 정수 타입의 메세지를 전달할것이다

	wg.Add(1)
	go square(&wg, ch) //새로운 고루틴을 만들고 동시에 실행한다.
	ch <- 9 //채널 인스턴스에 데이터를 할당한다
	wg.Wait() // 아래 선언되어 있는 함수가 위에서 생성했던 채널의 데이터를 빼서 계산하고 출력할 때까지 기다려 준다는 의미.
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <- ch
	time.Sleep(time.Second)
	fmt.Printf("Square: %d\n", n*n)
	wg.Done()
}