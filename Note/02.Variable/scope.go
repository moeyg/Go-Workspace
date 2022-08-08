// Variable Scope (변수 범위)

package main

import "fmt"

// var 키워드로 선언된 변수는 전역 변수로 사용이 가능
var integer int = 90

func main() {
	// 단축 키워드로 선언된 변수는 선언된 함수 내에서만 사용이 가능
	// 되도록 단축 키워드로 변수를 선언할 것을 권장
	pi := 3.14

	fmt.Println("integer : ", integer) // integer :  90
	fmt.Println("pi : ", pi) // pi :  3.14

	foo()
}

func foo() {
	fmt.Println("foo_integer : ", integer) // foo_integer :  90
	// fmt.Println("foo_pi : ", pi) // undefined: pi
}