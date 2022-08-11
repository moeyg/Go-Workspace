// Multiple Return Value : 다중 반환값

package main

import "fmt"

func Divide(a, b int) (int, bool) {
	if b == 0 {
		// return 값을 여러개 가능
		return 0, false
	}
	// return 값을 여러개 가능
	return a / b, true
}

// main 함수
func multipleReturn() {
	result, boolean := Divide(90, 3)
	fmt.Println(result, boolean) // 30 true

	result, boolean = Divide(90, 4)
	fmt.Println(result, boolean) // 22 true

	result, boolean = Divide(3, 0)
	fmt.Println(result, boolean) // 0 false
}