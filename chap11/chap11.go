package main

// chap11 Go 컬렉션1(배열)
func main() {
	// Go 컬렉션 중 배열은 java, javascript 와 동일하나
	// 문법적으로 차이가 있다.
	// java: int[] arr = {1, 2, 3}
	// javascript: var arr = [1, 2, 3] | var arr = new Array(1, 2, 3)
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	for _, s := range arr {
		println(s)
	}
	// 배열을 선언할 때 초기값을 설정하는 방법도 있다.
	// 만약 [] 안에 숫자가 아닌 ... 이 들어가서 배열크기를 생략하면
	// 자동으로 초기화 요소의 숫자만큼 배열 크기가 정해진다.
	var arr2 = [3]int{4, 5, 6}
	var arr3 = [...]int{7, 8, 9}
	for _, s := range arr2 {
		println(s)
	}
	for _, s := range arr3 {
		println(s)
	}

	// Go 의 경우 다차원 배열을 지원한다.
	// 다차원 배열의 경우 배열 크기 부분을 복수개로 지정하면 된다.
	var multiArr [2][3][4]int
	multiArr[0][1][2] = 10
	println(multiArr[0][1][2])
	var multiArr2 = [2][3]int{
		{11, 12, 13},
		{14, 15, 16}, // 끝에 콤마를 추가해야 컴파일 에러가 안생긴단.(Go 의 문법적 특징)
	}
	println(multiArr2[1][1])
}
