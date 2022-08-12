// Code Value : const로 선언한 상수명을 파라미터 코드 값으로 사용할 수 있다.

package main

import "fmt"

// 상수 선언
const (
	Jan = iota + 1
	Feb
	Mar
)

// 변수 선언
var may int = 5

func printMonth(month int) int {
	if month == Jan {
		fmt.Printf("%d월\n", month)
	} else if month == Feb {
		fmt.Printf("%d월\n", month)
	} else if month == Mar {
		fmt.Printf("%d월\n", month)
	} else {
		fmt.Print("??\n")
	}
	return month
}

func codeValue() {
	// 파라미터에 int 값 대신 const 상수명을 작성할 수 있다.
	printMonth(Jan) // 1월
	printMonth(Feb) // 2월
	printMonth(Mar) // 3월
	printMonth(0) // ??

	// 파라미터에 int 값을 넣어도 같은 값이 출력
	printMonth(1) // 1월
	printMonth(2) // 2월
	printMonth(3) // 3월
	printMonth(0) // ??

	// var로 선언한 경우는 파라미터에 변수명을 넣을 수 없다.
	printMonth(may) // ??
}