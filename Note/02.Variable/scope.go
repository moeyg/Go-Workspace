// Variable Scope (변수 범위)

package main

import "fmt"

// Package Variable : 전역 변수
// var 키워드로 선언된 변수는 전역 변수로 사용이 가능
var packageVariable string = "Package Variable"
var integer int = 90

func main() {
	// Local Variable : 지역 변수
	var localVariable string = "Local Variable"

	// 단축 키워드로 선언된 변수는 선언된 함수 내에서만 사용이 가능
	// 되도록 단축 키워드로 변수를 선언할 것을 권장
	pi := 3.14

	fmt.Println(packageVariable) // Package Variable
	fmt.Println(localVariable) // Local Variable

	fmt.Println("integer : ", integer) // integer :  90
	fmt.Println("pi : ", pi) // pi :  3.14

	foo()
}

func foo() {
	fmt.Println("foo_integer : ", integer) // foo_integer :  90
	// fmt.Println("foo_pi : ", pi) // undefined: pi
}