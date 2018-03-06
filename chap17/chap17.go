package main

import (
	"fmt"
)

// chap17 구조체1
func main() {
	// Go 의 특징적인 부분으로 구조체가 있다.
	// 구조체는 변수 여러개를 담을 수 있다.
	// 구조체 선언 방법은 일반변수를 선언하는 것과 동일하다.
	var a Rectangle // 구조체의 각 필드는 자료형의 기본값(문자열: "", 정수형: 0, 실수형: 0.0)으로 초기화 된다.
	fmt.Println(a)  // {0 0}
	// 구조체를 지역 변수 선언 형태가 아닌 포인터를 통해 메모리에 공간 할당을 할수도 있다.
	var b *Rectangle    // 구조체에 포인터 선언
	b = new(Rectangle)  // 구조체에 메모리 할당
	c := new(Rectangle) // 구조체 포인터 선언과 동시에 메모리 할당
	fmt.Println(b)      // &{0 0} 포인터이기 때문에 & 붙었다
	fmt.Println(c)      // &{0 0} 포인터이기 때문에 & 붙었다

	// 구조체는 인스턴스를 생성할 때 값을 초기화할 수 있다.
	// 변수명 = 구조체타입{}
	var d Rectangle = Rectangle{10, 20}
	fmt.Println(d) // {10 20}
	e := Rectangle{30, 40}
	fmt.Println(e)                        // {30 40}
	f := Rectangle{width: 50, height: 60} // 구조체 필드를 지정하여 값을 대입할 수 도 있다.
	fmt.Println(f)                        // {50 60}

	// 구조체 인스턴스의 값에 접근할 때는 json 형태와 비슷하다.
	// 구조체.필드를 통해 접근할수 있다.
	var g Rectangle
	var h *Rectangle = new(Rectangle) // 포인터 선언과 동시에 메모리 할당
	// . 을 통해 구조체 필드에 값 대입
	g.width = 100
	g.height = 200
	h.width = 1920
	h.height = 1024
	fmt.Println(g, h) // {100 200} &{1920 1024}

	// new 함수로 구조체의 메모리를 할당하는 동시에 값을 초기화하는 방법은 없다.
	// 하지만 다른 언어의 생성자(constructer) 를 흉내낼 수는 있다.
	i := NewRectangle(300, 400)
	fmt.Println(i) // &{300 400}
	// 위의 코드를 줄이면 다음과 같다.
	k := &Rectangle{200, 300} // 구조체를 초기화한 뒤 메모리 주소를 대입
	fmt.Println(k)            // &{200 300}

	// 구조체를 매개 변수로도 사용이 가능하다.
	rectScale1(i, 13)
	fmt.Println(i) // &{3900 5200} i 의 값이 영구적으로 바뀜
	rectScale2(g, 3)
	fmt.Println(g) // {100 200} g 의 값이 바뀌지 않음
}

// 구조체 문법 구조 type 구조체 명 struct {}
type Rectangle struct {
	width, height int // 같은 자료형끼리는 한줄로 표현할 수 있다.
}

// 구조체로 생성자 흉내내기
func NewRectangle(a, b int) *Rectangle {
	return &Rectangle{width: a, height: b}
}

// 구조체 매개변수 사용
func rectScale1(a *Rectangle, scale int) { // 매개변수에 포인터 사용으로 원래의 값에 영향을 미침
	a.width = a.width * scale
	a.height = a.height * scale
}
func rectScale2(a Rectangle, scale int) { // 매개변수에 인스턴스 사용으로 원래의 값에 영향이 없음
	a.width = a.width * scale
	a.height = a.height * scale
}
