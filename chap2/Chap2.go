package main

import "fmt"

// chap2 변수
func main() {
	// 변수 선언 방법1
	// var 를 이용하여 변수 선언 Go 의 경우 java 와 달리 변수 타입을 선언할 경우
	// 변수명 뒤에 변수 타입을 입력한다.
	var var1 int = 0
	fmt.Println(var1)
	// 변수 선언 방법2
	// Go 의 경우 변수 타입을 따로 선언하지 않아도 된다.
	var var2 = "hello!"
	fmt.Println("var2 --> " + var2)
	// 변수 선언 방법3
	// := (Short variable declaration(짧은 변수 선언)) 을 사용하여 변수를 선언 할 수 있다.
	// 이 경우 "var" 와 변수 타입을 지정하지 않아도 된다.
	// ※ 주의사항! := 오퍼레이터의 경우 func 밖에선 쓸 수 없다.!
	var3 := 0.3
	fmt.Println(var3)
	// 변수 선언 방법4
	// Go 에서는 여러개의 변수를 한번에 선언이 가능하다.
	// example1
	var i, j, k int
	// example2
	i, j, k = 1, 2, 3
	fmt.Println(i, j, k)
	// example3
	var r, t, y = "hi", 4, 5
	fmt.Println(r, t, y)
	// example4
	u, o, p := 10, 11, 223
	fmt.Println(u, o, p)
	// example5
	z, x := "i am the king!", 0.78
	fmt.Println(z, x)
	// example6
	var (
		a = "open source hurray!"
		b = "Golang is magic!"
	)
	fmt.Println(a, b)
}
