package main

import (
	"fmt"
)

// chap4 데이터 타입
func main() {
	// Go 의 최대의 단점은 generic 을 아직 완전히 지원하지 않는다는 것이다. (1.10 사용중...)
	// generic 을 완전히 지원하지 않고 있기 때문에 데이터 타입을 어느정도 숙지하고 있어야 한다.
	// 2018.02.27 기준 Go 개발진들이 generic 을 추가할 가능성을 언급했다!!

	// 1. 문자열
	// string: 문자열 타입인 string 은 한번 선언하면 변경할 수 없는 immutable 한 속성을 가지고 있다.
	// 2. 정수형
	// int, int8, int16, int32, int64: 부호 있는 정수
	// uint, uint8, uint16, uint32, uint64, uintptr: 부호 없는 정수 (uintptr 의 경우 포인터를 저장할 때 사용한다.)
	// 3. 실수형
	// float32, float64: 부동소수점
	// complex64, complex128: 실수부와 허수부로 구성된 복소수
	// 4. 기타
	// byte: uint8 과 동일하며 바이트 코드에 사용한다.
	// rune: int32 와 동일하며 유니코드 코드 포인트에 사용한다.

	// 문자열 에서 back_quota `` 와 double quota ""
	// back_quota `` 의 경우 raw_literal_line 이라고 하여 Escape 문을 문자열을 그대로 출력한다.
	rawLiteral := `아리랑\n아리랑\n아라리요` // \n 의 경우 NewLine 이 아니라 \n 문자열로 인식한다.
	fmt.Println(rawLiteral)
	// double quota "" 의 경우 inter_literal_line 이라고 하여 문자열의 Escape 문에
	// 특별한 의미를 부여하여 출력한다.
	interLiteral := "아리랑\n아리랑\n아라리요"
	fmt.Println(interLiteral) // \n 에 NewLine 의 의미를 부여한다.

	// 데이터 타입 변환
	// 다른 데이터 타입으로 변환하는 경우 DataType(value) 를 사용하여 변환한다.
	var i int = 100
	var k = uint(i)
	var j = float32(k)
	fmt.Println(i, k, j)

	str := "123123"
	bytes := []byte(str)
	str2 := string(bytes)
	fmt.Println(str, bytes, str2)

}
