/*
validator package

validation 태그를 써서 구조체, 필드 검증
슬라이스, 배열, 맵의 요소 하나씩 검증 가능하도록 지원
인터페이스에 대해서도 검증가능
검증 시 에러 메세지를 커스터마이즈 할 수 있음.
gin 웹 프레임워크에 적용할 수 있는 기본 검증모듈이다.
*/

package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

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
		Age int `validate:"-"` //무시
		//Name      string     `validate:"jack|joe"`                      //두 값 중 하나와 동일해야 함
		Subjects  Test0      `validate:"structonly"`                    //구조체가 중첩되어 있을 떄 이 구조체 자체가 할당이 되었는지까지만 확인한다. 이 구조체 내부의 필드까지는 검증안한다
		NickName  string     `validate:"omitempty"`                     //값이 기본값일 경우 검증하지 않음
		Cell      string     `validate:"required"`                      //기본값이 아니어야 한다
		Addresses [][]string `validate:"gt=0,dive,len=1,dive,required"` //배열, 슬라이스, 맴의 각 요소에 대해서도 유효성 검사흫 진행하는 dive 옵션

	}

	val := validator.New()

	t1 := Test1{}

	error := val.Struct(t1)
	if error != nil {
		fmt.Println(error)
	}
}

//기본적으로 http 엔드포인트 만들고 컨텍스트를 받아서 유효성 검사 함수를 걸어둘 수 있다.
/*
// CreateNewProject func for create a new project.
func CreateNewProject(c *fiber.Ctx) error {

    // ...

    // Create a new validator, using helper function.
    validate := utilities.NewValidator()

    // Validate all incomming fields for rules in Project struct.
    if err := validate.Struct(project); err != nil {
        // Returning error in JSON format with status code 400 (Bad Request).
        return utilities.CheckForValidationError(
            c, err, fiber.StatusBadRequest, "project",
        )
    }

    // ...
}

func CheckForValidationError(ctx *fiber.Ctx, errFunc error, statusCode int, object string) error {
    if errFunc != nil {
        return ctx.JSON(&fiber.Map{
            "status": statusCode,
            "msg":    fmt.Sprintf("validation errors for the %s fields", object),
            "fields": ValidatorErrors(errFunc),
        })
    }
    return nil
}
*/
