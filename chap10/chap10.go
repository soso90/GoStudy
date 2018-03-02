package main

// chap10 클로저(Closure)
func main() {
	// Go 에서 함수는 클로저(Closure) 로서 사용될 수 있다.
	/*
	*  클로저(Closure) 란?
	*  함수 밖에 있는 변수를 참조하는 함수 값을 일컫는다.
	*  이때의 함수는 밖의 변수를 함수 안에서 읽거나 쓸 수 있다.
	 */
	next := nextVal()
	println(next())
	println(next())
	println(next())
	// 다시 시작
	next = nextVal()
	println(next())
	println(next())
	println(next())
}

func nextVal() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
