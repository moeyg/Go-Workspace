// 1. Create your own type. Have the underlying type be an int.

// 2. Create VARIABLE of your new TYPE with the IDENTIFIER "num" using the "VAR" keyword

// 3. In func ex04

// 🅐 print out the value of the variable "num"
// 🅑 print out the type of the variable "num"
// 🅒 assign 90 to the VARIABLE "num" using the "=" OPERATOR
// 🅓 print out the value of the variable "num"

package main

import "fmt"

type number int
var num number

func ex04() {
	fmt.Println(num) // 0
	fmt.Printf("%T\n", num) // main.number

	num = 90
	fmt.Println(num) // 90
	fmt.Printf("%T\n", num) // main.number
}