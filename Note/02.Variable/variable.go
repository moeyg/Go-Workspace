// Variable : 변수

package main

import "fmt"

func variable () {
	// var = variable, 변수 선언
	// int = 변수 타입, 정수 타입
	var num int = 10
	// sting = 뱐수 타입, 문자열 타입
	var message string = "Hello World"

	// 변수값 재할당
	num = 90
	message = "Hello Go"

	// 변수값 출력
	fmt.Println(message, num)
}