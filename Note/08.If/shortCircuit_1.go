// Short Circuit : 조건문에서 논리 연산자를 사용했을 때, 좌변이 true면 우변은 무시되고 결과가 출력

package main

import "fmt"

func shortCircuit_1() {
	num := 1

	// 좌변 false && 우변 true -> false
	if 1 < num && num > 10 {
		fmt.Println("&&")
	}

	// 좌변 true || 우변 false -> true
	if 1 <= num || num > 10 {
		fmt.Println("||")
	}
}