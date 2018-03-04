package main

import (
	"fmt"
	"os"
)

// chap13 지연 호출
func main() {
	// Go 에는 함수의 지연 호출이라는 기능이 있다.
	// 지연 호출이란 특정 함수를 위치와 상관없이 함수가 끝나기 직전에 실행하는 기능을 말한다.
	// java 의 try ~ catch 문의 finally 와 비슷하다.
	// defer 함수명
	// 아래 함수를 실행하면 원래 "world"  가 먼저 콘솔에 출력되야하지만
	// defer 뒤에 쓰였기 때문에 가장 마지막에 실행된다.
	defer world()
	hello()
	hello()
	hello()
	hello()
	// defer 뒤에 함수를 정의할 수도 있다.
	defer func() {
		fmt.Println("hurray!")
	}()
	// defer 의 자료구조는 스택(LIFO) 와 동일하기 때문에
	// 가장 마지막에 defer 호출한 함수가 가장 먼저 실행 된다.
	// defer 의 경우 닫기가 필요할 때 유용하게 쓰일 수 있습니다. (파일 열고 닫기 등등)
	readFile()
}

func hello() {
	fmt.Println("hello!")
}

func world() {
	fmt.Println("world")
}

// 파일 호출 예시
func readFile() {
	file, err := os.Open("hello.txt")
	defer file.Close() // 다른 처리가 끝나면 마지막에 처리

	if err != nil {
		fmt.Println(err)
		return // file.Close() 함수 호출
	}

	buf := make([]byte, 100)
	if _, err = file.Read(buf); err != nil {
		fmt.Println(err)
		return // file.Close() 함수 호출
	}

	fmt.Println(string(buf))
}
