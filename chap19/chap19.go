package main

import (
	"fmt"
)

// chap19 구조체3 : 구조체 임베딩(Embedding)
func main() {
	// go 는 클래스를 제공하지 않으므로 상속의 개념또한 지원하지 않는다.
	// 하지만 구조체를 이용하여 상속을 구현할 수 있다.'

	// 학생 구조체 안에 사람 구조체를 가지고 있으므로 greeting() 함수를 사용할수 있다. (Has-a)
	var s student
	s.p.greeting() // greeting!!

	// 선생 구조체 안에 필드명 없이 선언하면 구조체가 해당 타입을 포함하는 관계가 된다. (Is-a)
	var t teacher
	t.person.greeting() // greeting!!

	// 만약에 자식 구조체에서 부모 구조체와 같은 함수 명을 가진 함수가 있다면
	// 자식 구조체가 가진 함수를 호출한다. 즉, 맨 아래의 구조체의 위의 구조체들의 함수를 오버라이딩 한다.
	t.greeting() // greeting student!!
}

// 사람 구조체 선언
type person struct {
	name string
	age  int
}

// 학생 구조체 선언
// 구조체 안에 사람 구조체 객체 포함
type student struct {
	p      person
	school string
	grade  int
}

// 인사 함수 구현
func (p *person) greeting() {
	fmt.Println("greeting!!")
}

// 선생 구조체 선언
// 선생 구조체 안에 사람 구조체를 필드명 없이 선언
type teacher struct {
	person
	school string
	class  string
	grade  int
}

func (t *teacher) greeting() {
	fmt.Println("greeting student!!")
}
