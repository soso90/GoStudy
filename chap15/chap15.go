package main

import (
	"fmt"
)

// chap15 포인터
func main() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	// Go 에서는 C, C++ 과 마찬가지로 메모리 주소를 표현하는 포인터를 지원한다.
	var numPtr *int
	fmt.Println(numPtr) // nil
	// C와 달리 *(Asterisk 에스터리스크) 를 타입 앞에 붙입니다.
	// 빈 포인터 변수는 바로 사용할 수 없기 때문에 new 함수로 메모리를 할당해야 합니다.
	numPtr = new(int)
	fmt.Println(numPtr) // 0xc0420100b0
	// new 함수는 자료형 크기에 맞는 메모리 공간을 할당 합니다.
	// Go 의 경우 가비지 컬렉션을 지원하기 때문에 메모리 할당 후 해제하지 않아도 됩니다.
	// 포인터형 변수에 값을 대입하거나, 가져오려면 역참조(dereference) 를 해야 합니다.
	*numPtr = 1
	fmt.Println(*numPtr) // 1
	// 포인터형 변수에는 메모리만 대입이 되기 때문에 값을 대입하려면 역참조를 해야 합니다.
	// 일반변수의 참조(reference) 를 사용하면 포인터형 변수에 대입할 수 있습니다.
	var num int = 1
	numPtr = &num
	fmt.Println(numPtr) // 0xc0420540b8: numPtr 포인터 변수에 저장된 메모리 주소
	fmt.Println(&num)   // 0xc0420540b8: 참조로 구한 num 의 메모리 주소
	// *: 역참조(dereference), 포인터형 변수의 값을 가져오거나 대입할 때 사용
	// &: 참조(reference), 일반 변수의 메모리 변수를 가져온다
	// ※ 주의 ! Go 에서는 아래와 같이 포인터 연산이나 메모리 값의 직접 대입은 허용하지 않는다.
	// numPtr++              // 컴파일 에러: 포인터 연산은 허용하지 않음
	// numPtr = 0xc0420100b0 // 컴파일 에러: 직접 메모리 값 대입 허용하지 않음

	// 포인터형 변수를 통해 함수 밖의 변수값을 변경할 수 있다.
	i := 1
	hello1(i)
	fmt.Println("일반 변수 대입 --> ", i) // 함수 밖의 있는 변수 값은 변경되지 않는다.
	hello2(&i)
	fmt.Println("포인터형 변수 대입 --> ", i) // 메모리 주소를 참조 하기 때문에 함수 밖의 변수도 변경된다.
}

func hello1(n int) {
	n = 2
}

func hello2(n *int) {
	*n = 2
}
