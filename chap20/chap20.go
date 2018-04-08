package main

import (
	"fmt"
)

// chap20: 인터페이스1
func main() {
	// 인터페이스는 메소드의 집합이다. 인터페이스는 메소드 자체를 구현하지는 않는다.
	// type 인터페이스명 interface{}
	// 인터페이스는 다른 자료형과 동일하게 함수의 매개변수, 리턴 값 구조체의 필드로 사용이 가능하다.
	var h empty    // 인터페이스를 선언하는 방법은 변수 선언과 같다.
	fmt.Println(h) // 빈 인터페이스는 nil 로 출력

	var i myInt = 10
	var p Printer // 인터페이스 선언

	p = i     // i 를 인터페이스 p에 대입
	p.Print() // 10

	r := rectangle{10, 20} // rectangle 선언
	p = r                  // r 을 인터페이스 p 에 대입
	p.Print()              // 10 20

	p1 := Printer(i) // 인터페이스를 선언하면서 i 로 초기화
	p1.Print()       // 10

	p2 := Printer(r) // 인터페이스를 선언하면서 r 로 초기화
	p2.Print()       // 10 20

	// 배열 (슬라이스) 형태로도 인터페이스를 초기화할 수 있습니다.
	pArr := []Printer{i, r} // 슬라이스 형태로 인터페이스 초기화
	for index, _ := range pArr {
		pArr[index].Print() // 슬라이스를 순회하면서 Print 메소드 호출
	}
	for _, value := range pArr {
		value.Print() // 슬라이스를 순회하면서 Print 메소드 호출
	}
}

type empty interface {
}

type myInt int // int 형을 myInt로 정의

type Printer interface {
	Print()
}

func (i myInt) Print() {
	fmt.Println(i)
}

type rectangle struct {
	width, height int
}

func (r rectangle) Print() {
	fmt.Println(r.width, r.height)
}
