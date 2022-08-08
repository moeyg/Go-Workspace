// Standard Output Function (표준 출력 함수)

package main

// fmt(format) : 표준 입출력을 제공하는 Package
import "fmt"

func output() {
	var message string = "Hello World"

	// Print : 함수 입력값을 출력
	fmt.Print("Print : ", message)

	// Println : 함수 입력값을 출력하고 개행
	fmt.Println("Println : ", message)

	// Printf : format에 맞도록 입력값을 출력
	// formatter(서식문자) : %d(정수), %f(실수), %s(문자열), %v(데이터 타입에 맞춰 기본 형태로 출력), %T(테이터 타입)
	fmt.Printf("Printf : %s, %v", message, message)
}

// 출력 결과
// Print : Hello WorldPrintln :  Hello World
// Printf : Hello World, Hello World