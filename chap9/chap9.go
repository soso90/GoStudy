package main

// chap9 익명함수
func main() {
	// Go는 javascript 와 마찬가지로 변수에 익명함수를 지정할 수 있다.
	multiple := func(i int, j int) int {
		return i * j
	}
	println(multiple(12, 15))
	// calc 함수에 multiple 함수 전달
	result := calc(multiple, 12, 15)
	println(result)
	// 파리미터에 직접 함수를 지정할 수 있다.
	result = calc(func(x int, y int) int { return x + y }, 12, 15)
	println(result)

	// ※ 중요 !!
	// Go 에서 중요한 문법 중 하나로 type 문을 사용한 함수의 원형을 정의하는 방법이 있다.
	// type 문은 구조체(struct), 인터페이스(interface) 등 사용자 정의(?)타입 을 만드는데 사용한다.
	// 또한 type 문을 통해 함수의 원형을 간단하게 구현할 수 있다.
	result = calc2(multiple, 12, 15)
	println("type --> ", result)
	// 파리미터에 직접 함수를 지정할 수 있다.
	result = calc2(func(x int, y int) int { return x + y }, 12, 15)
	println("type --> ", result)
	// type 문을 통해 함수의 원형을 정의하고 타 메서드에 함수를 전달하고 리턴 받는 기능을
	// 타 언어에서는 흔히 델리게이트(delegate) 라고 부른다. Go 에서도 이러한 Delegate 기능을 제공한다.
}

// type 문으로 함수 원형 구현
type archTypeCalc func(int, int) int

// 파리미터에 함수 지정
func calc(f func(int, int) int, k int, l int) int {
	return f(k, l)
}

// 파리미터에 함수 지정
func calc2(f archTypeCalc, k int, l int) int {
	return f(k, l)
}
