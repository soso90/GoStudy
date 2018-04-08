package main

import (
	"fmt"
)

// chap21 인터페이스2 - 덕 타이핑(Duck typing)
func main() {
	// 각 값이나 인스턴스의 실제 타입은 상관하지 않고 구현된 메소드로만
	// 타입을 판단하는 방식을 덕 타이핑(Duck typing) 이라고한다.
	var d Duck
	var p Person
	inTheForest(d) // 인터페이스를 통하요 오리의 quack, feathers 메소드 호출
	inTheForest(p) // 인터페이스를 통하여 사람의 quack, feathers 메소드 호출

	// Quacker 인터페이스는 quack, feathers 함수를 정의하고 있다.
	// 그리고 오리와 사람 모두 quack, feathers 함수를 구현했다.
	// 실제로 사용할 때에는 inTheForest 함수에 Quacker 인터페이스를 매개변수로 받는다.
	// 여기에 오리든 사람이든 inTheForest 함수에 넣을 수 있으며 quack, feathers 메소드를 호출한다.
}

type Duck struct{} // 오리 구조체 정의

func (d Duck) quack() { // 오리의 quack 메소드 정의
	fmt.Println("Quack!") // 오리 울음 소리
}

func (d Duck) feathers() { // 오리의 feathers 메소드 정의
	fmt.Println("a Duck has white and gray feather!")
}

type Person struct{} // 사람 구조체 정의

func (p Person) quack() { // 사람의 quack 메소드 정의
	fmt.Println("a Person imitates Duck!")
}

func (p Person) feathers() { // 사람의 feathers 메소드 정의
	fmt.Println("a Person pick up feather on ground and show that!")
}

type Quacker interface {
	quack()
	feathers()
}

func inTheForest(q Quacker) {
	q.quack()    // Quacker 인터페이스로 quack 메소드 호출
	q.feathers() // Quacker 인터페이스로 feathers 메소드 호출
}
