// Short Circuit을 간과했을 때 문제점

package main

import "fmt"

var count int = 0

func IncreaseAndReturn() int {
	fmt.Println("현재 count :", count)
	count++
	return count
}

func shortCircuit_2() {
	// if false && IncreaseAndReturn() < 5 일 때, Short Circuit으로 인해 뒤의 조건이 무시
	if true && IncreaseAndReturn() < 5 {
		fmt.Println("증가된 count :", count)
	}
}