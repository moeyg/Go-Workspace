// Type Conversion : 타입 변환
// Go는 강타입 언어로 연산의 각 항목의 타입은 반드시 같아야 한다.
// 각 항목의 타입이 다를 시 에러

package main

import "fmt"

func conversion() {
	var integer int = 10
	var float float64 = 10.1
	// fmt.Println(integer + float) // 계산 X
	fmt.Println(integer + int(float)) // 20
	fmt.Println(float64(integer) + float) // 20.1

	var integer32 int32 = 32
	var integer64 int64 = 64
	// fmt.Println(integer64 + integer32) // 같은 타입이어도 사이즈 다르면 계산 X
	fmt.Println(int32(integer64) + integer32) // 96
	fmt.Println(integer64 + int64(integer32)) // 96
	fmt.Println(int(integer64) + int(integer32)) // 96

	num1 := 90
	num2 := 90.9
	fmt.Println(num1 + int(num2)) // 180
	fmt.Println(float64(num1) + num2) // 180.9
}