package main

import (
	"fmt"
	"time"
)

func main() {
	//시간 표기법에 대한 예시를 제공함
	fmt.Println(time.Kitchen)

	//기본 생성 방법 3가지
	t1 := time.Now()
	fmt.Println(t1)

	t2 := time.Date(2023, 5, 5, 22, 37, 30, 111, time.UTC) //맨 마지막 인자는 타임존 표현에 사용한다
	fmt.Println(t2)

	t3 := time.Unix(26327214, 0)
	fmt.Println(t3)

	//t에서 제공하는 여러 메소드 사용하기
	fmt.Println(t2.Year(), t2.Month(), t2.Day(), t2.Weekday(), t2.Unix()) //...

	//유용: 다양한 시간 포맷 구하기
	/*
		const (
		   ANSIC       = "Mon Jan _2 15:04:05 2006"
		   UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
		   RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
		   RFC822      = "02 Jan 06 15:04 MST"
		   RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
		   RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
		   RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
		   RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
		   RFC3339     = "2006-01-02T15:04:05Z07:00"
		   RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
		   Kitchen     = "3:04PM"
		   // Handy time stamps.
		   Stamp      = "Jan _2 15:04:05"
		   StampMilli = "Jan _2 15:04:05.000"
		   StampMicro = "Jan _2 15:04:05.000000"
		   StampNano  = "Jan _2 15:04:05.000000000"
		)
		내가 이 포맷처럼 형식을 바꾸고 싶다면 위 상수 중 하나를 가지고 와서 format에 인자로 전달한다
	*/
	fmt.Println("formatted:", t2.Format(time.Kitchen))
	fmt.Println("formatted:", t2.Format(time.RFC1123))

	/*위 상수값들을 사용해 문자열을 시간 타입으로 변환할 떄도 쓸수있다*/
	s := "Mon May _2 05:04:17 2023"
	t, err := time.Parse(time.ANSIC, s)
	if err != nil {
		fmt.Println("parse error: ", err)
	}
	fmt.Println("parsed:", t)

	//월 숫자 표기: JS에선 1월을 0부터 매기지만 GO에서는 1월은 1부터 시작함
	fmt.Println(time.January == time.Month(1))

	/*
		UTC 시간 문자열 => KST 기준으로 변경하기
	*/
	utcTime := "2019-01-12 00:00:00.000"
	layout := "2006-01-02 15:04:05.000"
	utc, _ := time.Parse(layout, utcTime)

	loc, _ := time.LoadLocation("Asia/Seoul") //위치정보 로드
	kst := utc.In(loc)                        //실제 변환

	fmt.Println(utc)
	fmt.Println(kst)

	//더 간단히
	kstTime, _ := time.ParseInLocation(layout, utcTime, loc)
	fmt.Println(kstTime)

	fmt.Println(kst.UTC()) //간단히 UTC로 변경하기

	//시간 연산: Add
	exampleTime := time.Date(2019, 1, 12, 0, 0, 0, 0, time.UTC)
	afterTenSec := exampleTime.Add(time.Second * 10)
	afterTenMin := exampleTime.Add(time.Minute * 10)
	afterTenHour := exampleTime.Add(time.Hour * 10)

	fmt.Println(afterTenSec)  // 2019-01-12 00:00:10 +0000 UTC
	fmt.Println(afterTenMin)  // 2019-01-12 00:10:00 +0000 UTC
	fmt.Println(afterTenHour) // 2019-01-12 10:00:00 +0000 UTC

	//시간 연산: Sub
	sub1 := afterTenMin.Sub(exampleTime)
	fmt.Println(sub1)

	//시간 연산: 비교하기
	fmt.Println(exampleTime.Before(afterTenMin))
	fmt.Println(exampleTime.After(afterTenMin))
	fmt.Println(exampleTime.Equal(afterTenMin))

	//현재 작업을 잠시 중단 후, 일정 시각 이후에 특정 작업을 시작하고 싶을 때
	time.Sleep(1 * time.Second)
	fmt.Println("printed after 1 second") //0으로 하면 즉시 다음 작업 시작

	//인터벌을 만들고 싶을 떄(메모리 누수를 고려해 인터벌 완료후엔 해당 인터벌을 제거하는 것도 구현)
	cnt := 0
	c := time.Tick(2 * time.Second)
	done := make(chan bool)

	go func() {
		// wait for user input(터미널에서 엔터를 치거나) to stop the ticker
		//fmt.Scanln()

		//코드적으로 원하는 조건을 만들어서 ticker를 중지시킨다.
		for {
			if cnt == 5 {
				done <- true
			}
		}
	}()

	for {
		select {
		case next := <-c:
			fmt.Println(next)
		case <-done:
			// exit the loop
			fmt.Println(fmt.Sprintf("cnt:%d,", cnt), "Stopping ticker...")
			return
		}
		cnt++
	}
}
