package main

import (
	"fmt"
)

// chap18 구조체2: 구조체에 메소드 연결
func main() {
	// Golang  은 클래스가 없는 대신 구조체에 메소드를 연결할 수 있다.
	// func (리시버명 *구조체타입) 함수명 리턴값 자료형()
	rect1 := Rectangle{10, 20}
	// 함수를 선언할 때 func 키워드와 함수명 사이에 리시버 변수를 추가하였다
	// 리시버 변수를 통해 현재 인스턴스에 접근이 가능하다.
	// 리시버 변수에는 구조체 타입에 포인터를 사용하여 현재 인스턴스의 값을 가져오거나 변경할 수 있다.
	// C++, JAVA 의 this 와 비슷
	// . 을 통해 구조체 인스턴스에 연결된 메소드를 호출한다.
	rect1.scaleA(2)
	fmt.Println("area1: ", rect1)

	rect2 := Rectangle{10, 20}
	rect2.scaleB(2)
	fmt.Println("area2: ", rect2)

	// 메소드를 작성할 때 구조체 인스턴스의 값을 변경한다면
	// 포인터로 받고 일반적인 상황에서는 리시버 변수를 값 형태로 받는다.
	// 리시버 변수를 사용하지 않는 경우에는 _ 으로 생략이 가능하다.
	// func (_ Rectangle) scaleA(factor int) {}
}

type Rectangle struct {
	width, height int
}

// 리시버 변수의 구조체 타입이 포인터인 경우
// 원래의 값이 변경됨
func (rect *Rectangle) scaleA(factor int) {
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}

// 리시버 변수의 구조체 타입이 일반 인스턴스의 경우
// 값이 복사되었으므로 값이 변하지 않음
func (rect Rectangle) scaleB(factor int) {
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}
