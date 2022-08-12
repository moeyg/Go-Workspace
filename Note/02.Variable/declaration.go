// Variable Declaration : 변수 선언

package main

import "fmt"

func declaration()  {
	// 일반적인 변수 선언
	var integer int = 10
	fmt.Println("integer :", integer) // 10

	// Zero Value
	var initInt int
	fmt.Println("initInt :", initInt) // 0
	var initFloat float64
	fmt.Println("initFloat :", initFloat) // 0.0
	var initBool bool
	fmt.Println("initBool :", initBool) // false
	var initStr string
	fmt.Println("initStr :", initStr) // "" (빈 문자열)

	// 타입 지정 X : 값에 따라 타입이 지정된다.
	var str = "Hello World"
	fmt.Println(str) // Hello World
	var pi = 3.14
	fmt.Println(pi) // 3.14

	// 선언 대입문 := : 값에 따라 타입이 지정된다.
	number := 90
	fmt.Println(number) // 90
	msg := "Hello Go"
	fmt.Println(msg) // Hello Go
} 