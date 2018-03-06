package main

import (
	"fmt"
)

// chap16 Go 컬렉션2(맵)
func main() {
	// Go 는 java 와 마찬가지로 기본 자료형으로 맵(map) 을 지원한다.
	// Go 에서 맵에 공간을 할당하기 위해서는 make 함수를 사용해야 한다.
	// var 변수명 map[키 자료형]값 자료형 = make(map[키 자료형]값 자료형)
	// var a map[string]int = make(map[string]int)
	var b = make(map[string]int) // 맵을 선언할 때 키워드와 자료형을 생략할 수 도 있다.
	// 일반 변수와 마찬가지로 := 를 통해 var 를 생략 가능하다.
	c := make(map[string]string)
	// 맵을 생성하면서 키와 값을 초기화 하려면 {}를 사용한다.
	a := map[string]string{"key1": "keyValue1", "key2": "keyValue2"}
	// 생성된 맵에 데이터를 넣으려면 [] 안에 키를 지정하고 값을 대입하면 된다.
	b["key1"] = 100
	b["key2"] = 200
	c["key1"] = "hello"
	c["key2"] = "world"
	fmt.Println(a) // map[key1:keyValue1 key2:keyValue2]
	fmt.Println(b) // map[key1:100 key2:200]
	fmt.Println(c) // map[key2:world key1:hello]
	// 맵에서 할당되지 않은 키를 조회하면 빈값을 출력한다.
	// 값의 자료형이 실수형이나 정수형이라면 0, 문자형이라면 "" 을 출력한다.
	fmt.Println(a["key3"]) // ""
	fmt.Println(b["key3"]) // 0
	// if 문에서 활용
	value, ok := a["key3"]
	fmt.Println(value, ok) // "", false
	if !ok {
		fmt.Println("key not found!")
	}
	// for 문을 통해 맵의 모든 데이터를 순회할 수 있다.
	// for 키, 값 := range 맵 {}
	for key, value := range a {
		fmt.Println(key, value)
	}
	// range 리턴 값으로 key 를 사용하고 싶지 않다면 _ 을 쓰면된다.
	for _, value := range b {
		fmt.Println(value)
	}
	// 맵에서 할당된 데이터를 삭제하는데에는 delete 함수를 사용한다.
	// delete(맵, 키)
	delete(c, "key1")
	fmt.Println(c) // map[key2:world]
	// 맵의 값 안에는 일반 자료형 뿐만 아니라 맵 자체나 배열, 슬라이스도 들어갈 수 있다.
	mp := map[string]map[string]float32{
		"key1": map[string]float32{
			"key1": 12.3,
			"key2": 13.4,
			"key3": 14.1241,
		},
		"key2": map[string]float32{
			"key1": 200.31312,
			"key2": 1512.412142,
			"key3": 11251.1241,
		},
	}
	fmt.Println(mp) // map[key1:map[key1:12.3 key2:13.4 key3:14.1241] key2:map[key1:200.31313 key2:1512.4121 key3:11251.124]]
	mp2 := map[string][]int{
		"key1": []int{1, 2, 3},
		"key2": []int{4, 5, 6},
	}
	fmt.Println(mp2) // map[key1:[1 2 3] key2:[4 5 6]]
}
