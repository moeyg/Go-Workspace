// package = 코드를 묶는 단위, go 코드가 속한 package를 나타냄
// main = Program Starting point Package
// package main = Program Starting point Package로 단 1개
// package main = main 함수를 가지고 있는 package
// golang은 package main 1개 + 여러 개의 package
package main

// import "fmt" : fmt package를 사용
// "fmt" : 기본적인 출력과 입력 포맷 관련 기능을 제공
// Print/Scan, Println/Scanln, Printf/Scanf, Errorf 제공
import "fmt"

// func = function : 함수 선언
// main = Program Starting point
func main () {
	// fmt.Println = fmt package 안에 포함된 함수(Println)를 사용
	// Println = ln(line, 문자열)을 출력하고 개행 -> 한 줄 출력
	fmt.Println("Hello World")
}
