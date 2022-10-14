package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Age int
}

func (s *Student) String() string {
	return fmt.Sprintf("Student Age: %d", s.Age)
}

func PrintAge(stringer Stringer){
	//추상 인터페이스에서 Studnet로 구체적인 타입 변환. 
	if s, ok := stringer.(*Student); ok {
		fmt.Println(s,"타입변환 오케.")
	} 
	
	fmt.Printf("Age: %d", s.Age)
}

func main(){
	s := &Student{15}
	PrintAge(s)
}