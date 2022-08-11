// Logical Operator : 논리 연산자

// && : AND 연산자, 양변이 모두 true면 true
// || : OR 연산자, 양변 중 하나라도 true면 true
// ! : NOT 연산자, true면 false를 반환하고 false면 true 반환

package main

import "fmt"

func logical() {
	a := 10
	b := 10
	c := 20
	d := 30

	fmt.Println(a < b && c < d) // false
	fmt.Println(a == b && c < d) // true
	fmt.Println(a < b || c > d) // false
	fmt.Println(a < b || c < d) // true
	fmt.Println(!(a == b)) // false
}