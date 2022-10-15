//인터페이스의 타입 변환
// package main

// import "fmt"

// type Stringer interface{
// 	Stringer() string
// }

// type Student struct {
// 	Age int
// }

// func (s *Student) Stringer() string{
// 	return fmt.Sprintf("student Age: %d", s.Age)
// }

// func PrintAge(stringer Stringer){
// 	s := stringer.(*Student)
// 	fmt.Printf("Age: %d", s.Age)
// }

// func main(){
// 	s := &Student{15}
// 	PrintAge(s)
// }

//다른 인터페이스로 타입 변환하기
package main

type Reader interface{
	Read()
}

type Closer interface {
	Close()
}

type File struct{} //인터페이스 메서드를 구현한 타입

func (f *File) Read() {}
func (f *File) Close() {} //이 라인이 없으면 런타임 에러 발생

func ReadFile(reader Reader){
	if c, ok := reader.(Closer); ok  { //두개의 변수 선언시, 두번쨰 변수에는 타입 변환 가능 여부를 불린값으로 할당하기 떄문에 런타임 에러가 발생하지 않는다
		c.Close() //타입 변환이 성공했으면, 메서드를 실행한다.
	}
}

func main() {
	file := &File{}
	ReadFile(file)
}