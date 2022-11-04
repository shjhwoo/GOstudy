package main

import (
	"fmt"
	"sync"
	"time"
)

// 11.04 금요일 일정 간격으로 오래 걸리는 일을 동시에 실행하고 싶을 떄.

func square(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)
	terminate := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-terminate:
			fmt.Println("terminate")
			wg.Done()
		case n := <-ch:
			fmt.Println(n * n)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go square(&wg, ch)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	wg.Wait()
}

// func main() {
// 	var messages chan string = make(chan string) //채널명 채널타입 = make(메세지 타입), 메세지는 고루틴끼리 주고받는 정보!

// 	//채널에 데이터 넣기. 왼쪽에는 채널 인스턴스 명을 기재한다. 데이터를 넣을 때는 <- 연산자를 사용한다
// 	messages <- "this is a message"

// 	//채널에서 데이터를 빼기. 화살표 방향은, 데이터가 담기는 그릇을 향한다
// 	var msg string = <- messages
// }

// func main() {
// 	var wg sync.WaitGroup
// 	ch := make(chan int) //나는 정수 타입의 메세지를 전달할것이다. 메세지를 쌓아두는 큐이다. 즉, 우편함과 같은 곳.

// 	wg.Add(1)// 작업의 갯수를 설정한다.
// 	go square(&wg, ch) //새로운 고루틴을 만들고 동시에 실행한다.
// 	ch <- 9 //채널 인스턴스에 데이터를 할당한다= 우편함에 9라는 메세지를 넣는다.
// 	wg.Wait() // 아래 선언되어 있는 함수가 위에서 생성했던 채널의 데이터를 빼서 계산하고 출력할 때까지 기다려 준다는 의미.
// }

// func square(wg *sync.WaitGroup, ch chan int) {
// 	n := <- ch //우편함에 있던 메세지를 빼온다
// 	time.Sleep(time.Second)
// 	fmt.Printf("Square: %d\n", n*n)
// 	wg.Done() //작업을 완료했으므로 대기 작업 목록을 1개 감소시킨다.
// }

// func main () {
// 	ch := make(chan int)
// 	ch <- 9
// 	fmt.Println("Never Print!")
// }

//위의 코드는 데이터를 빼가는 고루틴이 없기 때문에 프로그램이 강제로 종료된다.
//버퍼를 만들어 주면 다른 고루틴이 빼가기 직전까지는 데이터를 저장하면서 기다릴 수 있다.
//버퍼는 다음과 같이 만든다
// func main(){
// 	ch := make(chan int, 2)
// 	ch <- 9
// 	ch <- 11
// 	ch <- 14
// 	fmt.Println("Print out!")
// }

//채널에 들어온 데이터를 뺴오는 예제
// func square(wg *sync.WaitGroup, ch chan int) {
// 	for n := range ch {
// 		fmt.Printf("Square %d\n", n*n) //데이터를 빼낸다.
// 		time.Sleep(time.Second)
// 	}
// 	wg.Done()
// }

// func main() {
// 	var wg sync.WaitGroup

// 	wg.Add(1)
// 	ch := make(chan int)
// 	go square(&wg, ch)

// 	for i := 0; i < 10; i++ {
// 		ch <- i * 2
// 	}
// 	close(ch)//********************************(반드시 wait 직전에 설정. 아니면 deadlock 된다.)
// 	wg.Wait()
// }

//만약에 조건에 따라 여러 개의 채널에서 데이터를 받아와서, 여러 개의 일을 좀 더 효율적으로, 멀티스레드처럼 하고싶다면

// func square(wg *sync.WaitGroup, ch chan int, quit chan bool) {
// 	for {
// 		select {
// 		case n := <- ch:
// 			fmt.Printf("Square %d\n", n*n)
// 		case <- quit:
// 			wg.Done()
// 			return
// 		}
// 	}
// }

// func main(){
// 	var wg sync.WaitGroup
// 	ch := make(chan int)
// 	quit := make(chan bool)

// 	wg.Add(1)
// 	go square(&wg,ch,quit)

// 	for i := 0; i < 10; i ++ {
// 		ch <- i * 2
// 	}
// 	quit <- true
// 	wg.Wait()
// }
