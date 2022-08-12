// Untyped Constant : 타입 없는 상수

package main

import "fmt"

const (
	PI = 3.14
	FloatPI float64 = 3.14
)

func unTyped() {
	var x int = PI * 100 // 오류 발생 X
	fmt.Println(x) // 314

	var y int = int(FloatPI * 100) // 오류 발생 -> 타입 변환
	fmt.Println(y) // 314
}