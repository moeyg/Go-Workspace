// Comparison Operator : 비교 연산자

package main

import "fmt"

func comparison() {
	x := 10
	y := 20

	fmt.Println(x == x) // true
	fmt.Println(x == y) // false

	fmt.Println(x != x) // false
	fmt.Println(x != y) // true

	fmt.Println(x < y) // true
	fmt.Println(x > y) // false

	fmt.Println(x <= y) // true
	fmt.Println(x >= y) // false
}