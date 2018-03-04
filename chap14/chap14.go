package main

import (
	"fmt"
)

// chap14 패닉(panic) 과 복구(recover)
func main() {
	// Go 에서 Exception 상황이 발생해서 프로그램이 종료된 것을 패닉(panic) 이라고 부른다.
	// 패닉(panic) 의 예시
	// a := [...]int{1, 2, 3}
	// for i := 0; i < 4; i++ {
	// 	fmt.Println(a[i])
	// }
	// Go 에서는 위와 같은 런타임 에러뿐만아니라 panic() 함수를 통해 사용자가 직접 에러를 발생할 수 있다.
	// java 의 new throw Exception 과 비슷하다.
	// panic("Error!")
	// fmt.Println("hello world!") // panic 함수를 통해 에러발생으로 실행이 안된다.

	// Go 에서는 recover 함수를 통해 panic 상황이 발생했을 때 바로 종료되지 않고 예외처리를 할 수 있다.
	f()
	fmt.Println("hello world!")
	// recover 함수는  반드시 지연 함수(defer) 를 통해서만 사용이 가능한 점을 반드시 기억하자!!!
}

func f() {
	defer func() {
		// recover 함수는 지연 호출로 사용해야 한다.
		// panic 이 발생해도 프로그램을 종료하지 않는다.
		// panic 함수에서 설정한 에러 메세지를 가져온다.
		s := recover()

		fmt.Println(s)
	}()
	// panic("Error !!")
	a := [...]int{1, 2, 3}
	for i := 0; i < 4; i++ {
		fmt.Println(a[i])
	}
}
