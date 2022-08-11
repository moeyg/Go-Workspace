// Named Return Value : 명명된 반환값

package main

import "fmt"

// return 값에 이름을 지정할 수 있다. -> result, boolean
func Multiplication(a, b int) (result int, boolean bool) {
	if a == 0 || b == 0 {
		// return 값에 명명
		result = 0
		boolean = false
		return
	}
	// return 값에 명명
	result = a * b
	boolean = true
	return
}

// main 함수
func namedReturn() {
	result, boolean := Multiplication(90, 10)
	fmt.Println(result, boolean) // 900 true

	result, boolean = Multiplication(90, 0)
	fmt.Println(result, boolean) // 0 false

	result, boolean = Multiplication(0, 10)
	fmt.Println(result, boolean) // 0 false
}