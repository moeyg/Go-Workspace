// const : 상수

package main

import "fmt"

func constFunc() {
	const num = 90
	const greet = "Hello World"

	const (
		a = 10.01
		b = 20.02
	)

	fmt.Println(num) // 90
	fmt.Println(greet) // Hello World
	fmt.Println(a) // 10.01
	fmt.Println(b) // 20.02

	fmt.Printf("%T\n", num) // int
	fmt.Printf("%T\n", greet) // string
	fmt.Printf("%T\n", a) // float64
	fmt.Printf("%T\n", b) // float64
}