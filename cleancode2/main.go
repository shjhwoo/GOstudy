package main

//리스코프 치환 원칙이란, 어떤 인터페이스에 속한 구조체 타입에 소속된 메서드가 증명이 가능하다면(잘 작동한다면)
// 그 인터페이스에 속한 모든 다른 구조체 타입에 소속된 메서드도 동일하게 증명가능하도록 코드를 작성해야 함을 의미함
//즉 메서드 구현에서 예외를 두어서는 안된다 == 함수 계약 관계를 깨트리면 안된다
func main() {

}

//예를 들어 다음과 같이 코드를 작성하여야 리스코프 치환 원칙을 준수했다고 할 수 있다/

type Poster interface {
	SendReport() string
}

type AnaloguePost struct{}

func (a *AnaloguePost) SendReport() string { //반드시 문자열을 리턴해야 한다는 함수 계약 관계가 성립된다
	return "send with hand"
}

type DigitalPost struct{}

func (d *DigitalPost) SendReport() string { //리스코프 치환 원칙에 의해 마찬가지로 문자열을 반환해야 한다
	return "send with digital"
}

func Sender(p Poster) string {
	return p.SendReport()
}

//다음의 경우는 리스코프 치환 원칙을 위배한 경우이다.

func BadSender(p Poster) string {
	if _, ok := p.(*AnaloguePost); ok { //이 타입에 대해서만 예외를 두었기 때문이다.
		panic("Can't handle :(")
	}
	return p.SendReport()
}
