package main

import "fmt"

// chap5 연산자
func main() {
	// 산술 연산자: 사칙연산자(+-*%/), 증감연산자(++,--)
	a := (1 + 2) / 4 * 60
	fmt.Println("a -->", a)
	// 관계 연산자: == , !=, <=, >=
	b := 34
	if a != b {
		fmt.Println("wow!")
	}
	// 논리 연산자: &&, ||
	var c bool = true
	if c && (a != b) {
		fmt.Println("delicious!")
	}
	// Bitwise 연산자 &, |, <<, >> (AND, OR, XOR)
	// 아직 이게 뭔 연산자 인지 어떻게 쓰는 건진 잘 이해가 안된다.
	b = (a & b) << 2
	fmt.Println("b -->", b)
	// 할당 연산자 =, |=, *=, >>=
	a = 100
	fmt.Println("b -->", a)
	a |= 100
	fmt.Println("b -->", a)
	a *= 100
	fmt.Println("b -->", a)
	a >>= 100
	fmt.Println("b -->", a)
	// ** 포인터 연산자 ** & , * 을 사용하여 해당 변수의 주소를 참조하거나 이를 dereference 를 한다.
	k := 100
	fmt.Println("k -->", k)
	var j = &k
	fmt.Println("j -->", j)
}
