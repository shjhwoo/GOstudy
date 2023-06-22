/*
validator package

validation 태그를 써서 구조체, 필드 검증
슬라이스, 배열, 맵의 요소 하나씩 검증 가능하도록 지원
인터페이스에 대해서도 검증가능
검증 시 에러 메세지를 커스터마이즈 할 수 있음.
gin 웹 프레임워크에 적용할 수 있는 기본 검증모듈이다.
*/

package main

func main() {

	//검증함수는 인풋이 유효하지 않을 때 InvalidValidationError 리턴.
	//즉 에러가 nil인지 아닌지를 확인해야 한다.
	//혹은 에러가 nil인 경우 (거의 이럴 일은 없지만) 에러 타입이 InvalidValidationError인지 아닌지 한번 더 확인.
	// err := validate.Struct(mystruct)
	// validationErrors := err.(validator.ValidationErrors)

	//기본 사용
	type Test0 struct {
		Age int `validate:"max=10,min=1"` //주의: `validate:"min=10,max=0"` 처럼 쓰면 당연히 안됨
	}

	//여러개의 struct에 걸쳐 필드를 검증하려면 반드시 커스텀한 검증함수를 별도로 만들어야 하지만
	//그렇지 않다면 기본으로 제공하는 검증함수를 쓰면된다

	//쓸 수 있는 기본 검증 태그들.

	type Test1 struct {
		Age      int    `validate:"-"`          //무시
		Name     string `validate:"jack|joe"`   //두 값 중 하나와 동일해야 함
		Subjects Test0  `validate:"structonly"` //구조체가 중첩되어 있을 떄 이 구조체 자체가 할당이 되었는지까지만 확인한다. 이 구조체 내부의 필드까지는 검증안한다
		NickName string `validate:"omitempty"`  //값이 기본값일 경우 검증하지 않음
		Cell     string `validate:"required"`   //기본값이 아니어야 한다
		//...
	}
}
