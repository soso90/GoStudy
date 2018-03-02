package main

// chap8 함수
func main() {
	// Go 의 함수 사용은 기본적으로 javascript, java 와 같다.
	// 하지만 파라미터를 전달하는 방식이나 문법에서는 차이가 있다.

	// 차이점 예시
	// javascript, java : 함수명 (parameterType value)
	// Go : 함수명(value ParameterType) | 함수명(value ParameterType) (value returnType)

	// Go 에서 함수에 파라미터 인자 전달은 크게 2가지로 나뉜다.
	// 1. 값을 전달: 파라미터의 값을 그대로 복사해서 전달한다.
	text := "hurray!"
	printText1(text)
	// 2. 참조값 전달: 포인터를 사용하여 파라미터 값을
	// 그대로 복사하는 것이 아니라 메모리 영역 주소를 전달한다.
	// 변수 앞에 & 을 붙여 파라미터가 변수의 주소를 표시하게 된다.
	// text 주소에 데이터를 쓰기 위해서는 *text = "" 와 같이 앞에 * 을 붙이는데
	// 이를 흔히 Dereferencing 이라 한다.
	printText2(&text)

	// java 나 javascript 에서는 파라미터의 갯수를 가변적으로 사용할 수 없으나
	// Go에서는 ... 을 이용하여 파라미터를 가변적으로 사용이 가능하다.
	variableParam(text, "hello!", "world", "a cafela")
	// java 나 javascript 에서는 복수 개의 리턴 값이 필요할 때는
	// array, list, map, json 등을 사용하였지만
	// Go 에서는 리턴의 갯수를 복수 개로 지정할 수 있다.
	// 1. 리턴 타입만 지정
	rtn1, rtn2 := variableReturn1()
	println(rtn1, rtn2)
	// 2. 리턴 변수 지정
	rtn3, rtn4 := variableReturn2(1, 2, 3, 4)
	println(rtn3, rtn4)
}

func printText1(text string) {
	println(text)
}

func printText2(text *string) {
	println(text)
}

func variableParam(text ...string) {
	for _, s := range text {
		println(s)
	}
}

// 리턴 타입 만 지정 했을 때
func variableReturn1() (string, int) {
	rtn1 := "1year --> "
	rtn2 := 365
	return rtn1, rtn2
}

// 리턴 변수를 지정 했을 때
func variableReturn2(nums ...int) (count int, total int) {
	for _, s := range nums {
		total += s
	}
	count = len(nums)
	return
}
