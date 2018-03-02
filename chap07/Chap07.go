package main

import (
	"fmt"
)

// chap7 반복문
func main() {
	// java 와 마찬가지로 for 를 사용한다.
	sum := 0
	for i := 0; i < 20; i++ {
		sum++
	}
	fmt.Println(sum)
	// for 문 안에 증감식이 아닌 조건문을 넣음으로써
	// java while 문처럼 사용이 가능하다.
	sum = 0
	for sum < 100 {
		sum += 2
	}
	fmt.Println(sum)
	// for 문에 아무 조건 없이 쓰면 무한반복문이 되고 Ctrl + c 로 빠져나온다.
	// for {
	// 	fmt.Println("forever loop1")
	// }
	// Go 에서도 java 와 마찬가지로 배열 의 요소 만큼 반복문을 쓸 수 있다.
	// for 인덱스, 요소값 := range 컬렉션
	var arr = []string{"1", "2", "3", "4", "5"}
	for index, ele := range arr {
		fmt.Println(index, ele)
	}
	// for 문에서 continue, break, goto
	a := 0
	for a < 15 {
		if a == 10 {
			a += a
			continue
		}
		a++
		if a > 11 {
			break
		}
	}
	if a == 20 {
		goto L1
	} else {
		fmt.Println("Odd")
	}

L1:
	fmt.Println("Even")
}
