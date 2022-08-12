// Const : 상수, 값이 변하지 않는 수, 재할당 X
// 상수는 동적 메모리 공간을 사용하지 않고, literal 값으로 되어있다.
// 따라서 좌변으로 사용할 수 없다.

// 상수 선언
// const ConstName type = value

package main

import "fmt"

func constDeclaration() {
	const language = "Go"
	const num = 90

	const constPI float64 = 3.141592653489783228
	var varPI float64 = 3.141592653489783228

	fmt.Printf("%s\n", language) // Go
	fmt.Printf("%d\n", num) // 90
	fmt.Printf("const 원주율 : %f\n", constPI) // const 원주율 : 3.141593
	fmt.Printf("var 원주율 : %f\n", varPI) // var 원주율 : 3.141593

	// constPI = 3.14 -> const는 재할당 X
	varPI = 3.14 // var는 재할당 가능
	fmt.Printf("var 원주율 : %.2f\n", varPI) // var 원주율 : 3.14
}