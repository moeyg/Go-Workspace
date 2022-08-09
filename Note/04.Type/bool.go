// bool : true / false를 나타낼 수 있는 타입

package main

import "fmt"

func boolean() {
	var boolean bool
	fmt.Println(boolean) // false -> bool의 zero값은 false

	boolean = true
	fmt.Println(boolean) // true

	a := 10
	b := 20
	fmt.Println(a == b) // false
	fmt.Println(a != b) // true
	fmt.Println(a >= b) // false
	fmt.Println(a < b) // true
}