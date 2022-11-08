package main

func main() {}

//인터페이스 분리 원칙
//클라이언트는 자신이 이용하지 않는 메서드에 의존하지 않아야 한다
//인터페이스를 분리하면 불필요한 메서드들과 의존관계가 끊어져 더 가볍게 인터페이스를 이용할 수 있다

type Report interface {
	Report() string
	Pages() int
	Author() string
}

func SendReport(r Report) {
	send(r.Report())
}

//분리하여야 한다.

type Reporter interface {
	Report() string
}

type WrittenInfo interface {
	Pages() int
	Author() string
}

func SendReport2(r Report) {
	send(r.Report())
}
