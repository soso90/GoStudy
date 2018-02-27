package main

import "fmt"

// chap3 상수
func main() {
	// 상수 선언 방법
	// "const" 를 사용한다.
	// example1
	const i int = 0
	fmt.Println(i)
	// example2
	const j float32 = 0.3
	fmt.Println(j)
	// example3
	const k, l = 3, "상수!"
	fmt.Println(k, l)
	// example4
	const (
		z = "i am the king!"
		x = 13124
	)
	fmt.Println(z, x)
	// example5
	// ※ Go 에서는 iota (identifier) 를 이용해 변수, 상수의 값을 자동을 1씩 증가 시킬 수 있다!
	const (
		a = iota
		b
		c
	)
	fmt.Println(a, b, c)
}
