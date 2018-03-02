package main

import "fmt"

// chap12 Go 컬렉션2(Slice)
func main() {
	// Go 의 배열은 배열크기 안에 동일한 타입의 데이터를 연속으로 저장하지만
	// 배열의 크기를 동적으로 증가시키거나 부분 배열을 발췌하는 등의 기능을
	// 가지지 않고 있다. 그래서 Slice 는 배열에 기초하지만 배열 제약점들을
	// 넘어 개발자들에게 여러가지 편의를 제공한다.
	// Go 의 Slice 선언은 배열과 같지만 배열과 달리 크기를 지정하지 않는다.
	var arr []int
	arr = []int{1, 2, 3}
	println(arr[1])
	// Go 의 Slice 선언의 또다른 방법은 make() 내장함수를 사용하는 것이다.
	// make(타입, 길이, 용량(내부 배열의 최대 길이))
	arr2 := make([]int, 3, 5)     // 용량을 지정하면 모든 요소가 0 인 슬라이스를 만든다.
	println(len(arr2), cap(arr2)) // 내장함수 len(), cap() 각각 길이와 용량을 리턴한다.
	// slice 에 별도의 길이나 용량을 지정하지 않으면 리턴 값이 nil 인 Slice 를 만든다.
	var arr3 []int
	if arr3 == nil {
		println("arr3 is nil")
	}
	println(len(arr3), cap(arr3))
	// Go 의 Slice 는 앞서 말한대로 기본 배열과 달리 부분 요소 발췌도 가능하다.
	// Slice[처음 인덱스 : 원하는 인덱스 + 1]
	arr4 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := arr4[2:5]
	fmt.Println(s)
	arr4 = arr4[5:] // 마지막 인덱스를 생략하면 자동으로 Slice 의 마지막 인덱스가 대입된다.
	fmt.Println(arr4)
	// 배열은 기본적으로 크기 이상의 임의의 데이터를 요소에 추가할 수 없지만
	// Go 의 Slice 추가할 수 있다.
	// append(Slice, 추가할 값)
	arr4 = append(arr4, 11)
	arr4 = append(arr4, 11, 12, 13) // 복수 개를 추가할 때
	fmt.Println(arr4)
	// ※ 알고 있으면 좋은 것
	// append 함수의 프로세스를 살펴보면 내부적으로 다음과 같은 일이 일어난다.
	// Slice 의 용량 확인 -> 남아있으면 그 용량 내에서 길이를 변경하여 데이터를 추가
	//                    -> 남아 있지 않다면 용량을 2배로 하여 새로운 슬라이스를 만들고 기존 데이터를 복사
	// append 프로세스 확인
	append_process_chk()
	// append 함수를 이용하여 Slice 에 다른 Slice 를 병합할 수 있다.
	// append(Slice, 추가할 Slice...)
	arr = append(arr, arr4...)
	fmt.Println(arr)
	// Slice 는 copy 내장함수를 이용하여 다른 Slice 를 복사할 수 있다.
	// copy(Slice , 복사할 Slice)
	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2)
	copy(target, source)
	fmt.Println(target)               // [0 1 2 ] 출력
	println(len(target), cap(target)) // 3, 6 출력
}

func append_process_chk() {
	var slice1 = make([]int, 0, 3)
	for i := 0; i <= 21; i++ {
		slice1 = append(slice1, i)
		fmt.Println("slice1 --> ", len(slice1), cap(slice1))
	}
	fmt.Println(slice1)
}
