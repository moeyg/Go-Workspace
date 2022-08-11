// Function Definition : 함수 정의

package main

import "fmt"

func Add(a int, b int) int {
	// a와 b를 더한 결과를 반환
	return a + b
}

// func : 함수 정의 키워드
// Add : 함수명
// (a int, b int) : 매개변수
// int : 반환 타입
// {} : 함수 코드 블록

// main 함수
func definition() {
	// 함수 호출
	sum := Add(3, 4)
	fmt.Println(sum)
}