// Decimal Output (정수 출력) : %d
// fmt.Printf("%d", int)

package main

import "fmt"

func decimal() {
	var integer int = 90

	// 5칸을 오른쪽 정렬
	fmt.Printf("%5d\n", integer) //    90

	// 5칸을 오른쪽 정렬 + 공백을 0으로
	fmt.Printf("%05d\n", integer) // 00090

	// 5칸을 왼쪽 정렬
	fmt.Printf("%-5d", integer) // 90   
}