package main

import "fmt"

func main() {}

//의존 관계 역전 원칙
//상위계층이 하위계층에 의존하는 관계를 역전시킨다.
/*
예를 들어, 키보드와 네트워크 모듈에 의존하고 있는 전송 모듈을 구현한다고 가정하자.
전송 모듈은 키보드로부터 입력을 받고, 네트워크에 이를 전송해준다.
하지만 해당 모듈은 키보드와 네트워크, 즉 하위 계층에 강하게 결합되어 있기 때문에 입력 장치가 변경될 경우 코드 재사용성이 떨어진다는 문제가 있다.
이를 추상화하면 다음과 같다.
키보드는 "입력", 네트워크는 "출력" 모듈로 추상화하고, 각각 추상화된 모듈끼리만 전송 모듈로 연결해주는 것이다.
즉 전송 모듈은 입력과 출력이라는 추상화된 모듈에만 의존하기 때문에 어떤 입력, 출력장치를 도입하더라도 코드를 그대로 재사용할 수 있게 된다.
*/

/*
두번째 원칙
추상 모듈은 구체화된 모듈에 의존해서는 안된다. 구체화된 모듈은 추상 모듈에 의존해야한다
예를 들면 다음과 같은 코드는 의존 관계 역전 원칙을 위배한 경우이다.
*/

// type Mail struct{
// 	alarm Alarm
// }

// type Alarm struct {}

// func (m *Mail) OnRecv(){
// 	m.alarm.Alarm()
// }

//의존관계 역전원칙을 지키도록 코드를 작성한다면, 메일 수신은 이벤트, 알람은 이벤트리스너라는 모듈로 추상화할 수 있을것이다!

type Event interface {
	Register(EventListener)
}

type EventListener interface {
	OnFire()
}

type Mail struct {
	listener EventListener
}

func (m *Mail) Register(eventListener EventListener) {
	m.listener = eventListener
}

func (m *Mail) OnRecv() {
	m.listener.OnFire()
}

type Alarm struct {
}

func (a *Alarm) OnFire() {
	fmt.Println("메일 받아")
}

var mail = &Mail{}
var listener EventListener = &Alarm{}

mail.Register(listener)
mail.OnRecv()


//나쁜 코드 개선하기

// type Player struct {
// 	name string
// }

// type Monster struct {
// 	hp int
// }

// func (p *Player) Name() string {
// 	return p.name
// }

// func (p *Player) Attack(m *Monster){
// 	m.DealDamage(p, 100)
// }

// func (m *Monster) DealDamage(attacker *Player, damage int) {
// 	m.hp -= damage
// 	if m.hp < 0 {
// 		fmt.Println(attacker.name(), "가 날 죽였다")
// 	}
// }

//=======

type Player interface {
	Name() string //attack을 안 넣은 이유는 한 인터페이스에 메서드 하나만 넣어서 인터페이스를 분리하기 위함. 또한, 공격 말고도 다른 메서드도 구현을 할려고 이래둔 것 같다. 
}

type Monster interface {
	DealDamage(attacker *Player, damage int)
}

type p1 struct {
	name string
}

func (p *p1) Name() string { return p.name }
func (p *p1) Attack(m *Monster) {
	m.DealDamage(p, 100)
}


type m1 struct {
	hp int
}

func (m *Monster) DealDamage(attacker *Player, damage){
	m.hp -= damage
	if m.hp < 0 {
		fmt.Println(attacker.name(),"가 날 죽였다")
	}
}

