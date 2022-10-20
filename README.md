# GO

```
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var wg sync.WaitGroup

// func sumAtoB(a,b int){
// 	sum := 0
// 	for i:=0;i<a;i++{
// 		sum += i
// 	}
// 	fmt.Println(sum)
// 	wg.Done()
// }

// func main() {

// 	wg.Add(10)
// 	for i := 0; i < 10; i++ {
// 		go sumAtoB(1, 500)
// 	}
// }

/*
멀티프로세스: 여러개의 실행중인 프로그램 목록
멀티스레드: 하나의 프로그램안에 있는 실행 흐름이 여러 개 있는 것
고루틴: 경량스레드.

아무리 많이 고루틴을 만들어도 컨텍스트 스위칭 비용이 발생안함
*/

//데드락 문제 해결 위한 뮤텍스 사용
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var mutex sync.Mutex

// type Account struct {
// 	Balance int
// }

// func main() {
// 	var wg sync.WaitGroup

// 	account := &Account{10}
// 	wg.Add(10)
// 	for i := 0; i < 10; i++ {
// 		go func() {
// 			for {
// 				DepositeAndWithdraw(account)
// 			}
// 		}()
// 	}
// }

// func DepositeAndWithdraw(account *Account) {
// 	mutex.Lock()
// 	defer mutex.Unlock()
// 	if account.Balance < 0 {
// 		panic(fmt.Sprintf("ddd"))
// 	}

//  }

//뮤텍스 문제 해결: 작업 영역 분핧 또는 고루틴 간의 간섭 삭제

package main

import (
	"fmt"
	"sync"
	"time"
) 

type Job interface {
	Do()
}

type SquareJob struct {
	index int
}

func (j * SquareJob) Do() {
	fmt.Printf("%d 작업 시작", j.index)
	time.Sleep(1* time.Second)
	fmt.Printf("%d 작업 완료 -결과: %d\n", j.index, j.index * j.index)
}

func main (){
	var joblist [10]Job
	for i := 0; i < 10; i++ {
		joblist[i] = & SquareJob{i}
	}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0	; i < 10; i++ {
		job := joblist[i]
		go func(){
			job.Do()
			wg.Done()
		}() //선언 후 즉시실행함수
	}
	wg.Wait() //모든 고루틴 완료될 때까지 대기
}
```
