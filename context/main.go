package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// var wg sync.WaitGroup

// func main() {
// 	wg.Add(1)
// 	ctx, cancel := context.WithCancel(context.Background()) // 작업시간을 지정하고 싶은 경우: context.WithTimeout(context.Background(), time.Second*3). 똑같은 효과를가진다
// 	go printEverySecond(ctx)
// 	time.Sleep(5 * time.Second)
// 	cancel()

// 	wg.Wait()
// }

// func printEverySecond(ctx context.Context) {
// 	tick := time.Tick(time.Second)
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			wg.Done()
// 			return
// 		case <-tick:
// 			fmt.Println("tick")
// 		}
// 	}
// }

//컨텍스트에 값을 전달할 수도 있다
//별도의 조건이나 지시사항을 작업자에게 전달하는 기능이다
//이때는 context.WithValue를 사용한다ㅏ

// var wg sync.WaitGroup

// func main() {
// 	wg.Add(1)
// 	ctx := context.WithValue(context.Background(), "number", 9) //항상 상위 컨텍스트를 인수로 전달해줘야 한다 어떤 경우든.
// 	go square(ctx)
// 	wg.Wait()
// }

// func square(ctx context.Context) {
// 	if v := ctx.Value("number"); v != nil {
// 		n := v.(int) //빈 인터페이스이므로 타입을 반드시 변환해야 한다
// 		fmt.Printf("squaare:%d", n*n)
// 	}
// 	wg.Done()
// }

//컨텍스트 연이어 전달하여 여러 조건을 가지는 작업을 지시할 수 있게된다ㅏ

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "number", 9)
	ctx = context.WithValue(ctx, "keyword", "rabbit")
	go work(ctx)
	time.Sleep(time.Second * 10)
	cancel()
	wg.Wait()
}

func work(ctx context.Context) {
	tick := time.Tick(time.Second * 2)
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		case <-tick:
			if v := ctx.Value("number"); v != nil {
				fmt.Println("tick", (v.(int)))
			}
			if s := ctx.Value("keyword"); s != nil {
				fmt.Println("tick " + (s.(string)))
			}
		}
	}
}
