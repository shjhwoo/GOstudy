package main

import "fmt"

func main() {
	// var array [10]int
	// var slice []int = array[1:3:5]
	// fmt.Println(slice)
}

//단일 책임 원칙을 지켜서  설계해라. 코드 재사용성을 높여준다

type Report interface {
	Report() string
}

type financeReport struct {
	report string
}

func (f financeReport) Report() string {
	return f.report
}

type reportSender struct {
}

func (r reportSender) SendReport(report Report) {
	report.Report()
}

//어떤 종류의 보고서이든지 상관없이 보고서 전송방식이 동일하다면, 보고서들을 모두 하나의 인터페이스에 종속되게 만들고
//그 인터페이스를 인수로 하는 보고서 전송 함수를 만들게 되면 각 보고서생성 함수마다 보고서 전송 함수를 일일이 선언할 필요가 없으므로 편리해진다

// ======개방 폐쇄 원칙
// 확장에는 열려 있고 변경에는 닫혀있다. 즉 기존의 코드는 수정하면 안되지만 새로운 기능은 어렵지않게 추가할 수 있어야 한다
func SendReport(r *Report, method SendType, receiver string) {
	switch method {
	case Email:
	case Post:
		BadStmt

		//위와 같이 작성할 경우 새로운 전송방식이 생길때마다 기존의 함수를 변경해야한다

		type reportSender interface {
			Send(r *Report)
		}

		type EmailSender struct {
		}

		(func(e *EmailSender) Send)(r * Report)

		type FaxSender struct{}

		(func(e *FaxSender) Send)(r * Report)
	}
}
