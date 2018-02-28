package main

import (
	"fmt"
)

// chap 6 조건문
func main() {
	// Go 의 조건문은 조건절을 () 안에 넣지 않고
	// if 바로 옆에 그대로 입력한뒤 {} 로 감싼다.
	var k bool = false
	if k != true {
		fmt.Println("hurray!")
	}
	// if, else if , else 의 사용법의 경우 java 와 동일하다.
	i := 13
	if i > 13 {
		fmt.Println("Not More!")
	} else if i < 0 {
		fmt.Println("Not Enough")
	} else {
		fmt.Println("Etc...")
	}
	// Go 에서는 if 조건절에 간단한 문장을 같이 실행할 수 있다.
	if j := i; j == 13 {
		fmt.Println("Go Hurray!")
	}
	// switch 문
	company_no := 400
	name := ""
	switch company_no {
	case 0:
		name = "Google!"
	case 13, 15:
		name = "IBM!"
	case 27:
		name = "Microsoft!"
	case 400:
		name = "LG!"
	case 567:
		name = "Bethesda Studio1"
	case 627:
		name = "Code1System1"
	default:
		name = "Samsung!"
	}
	fmt.Println(name)
	switch x := company_no << 213; x - 1 {
	case 0:
		name = "IBK1"
	case 1, 2, 3:
		name = "Nexon!"
	default:
		name = "INZEN!"
	}
	fmt.Println(name)
	// java 의 switch 문과 Go 의 switch 문의 차이점은
	// switch 문 뒤에 expression 를 true 로 간주하고 진행하는 것이다.
	// 또한 복잡한 expression 을 사용할 수 있다.
	// 다른 특징적인 것으로는 switch 문으로 변수의 타입을 골라낼 수 있다.
	name = doSwitch(company_no, name)
	fmt.Println(name)
	// Go 의 경우 break 가 따로 없어도 case 문을 타면 자동으로 break 를 실행한다.
	// 만약 break 를 안하고 계속 밑의 조건문을 타려면 fallthrough 를 사용한다.
	check(23)

}

func doSwitch(i interface{}, name string) string {
	switch i.(type) {
	case int:
		name = "INT"
	case float32:
		name = "FLOAT32"
	case uint:
		name = "UINT"
	case string:
		name = "STRING"
	default:
		name = "I don`t Know variable type"
	}
	return name
}

func check(val int) {
	switch val {
	case 1:
		fmt.Println("1 이하")
		fallthrough
	case 2:
		fmt.Println("2 이하")
		fallthrough
	case 3:
		fmt.Println("3 이하")
		fallthrough
	default:
		fmt.Println("default 도달")
	}
}
