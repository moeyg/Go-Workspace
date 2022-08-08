// 1. At the package level scope, using the "var" keyword, create a VARIABLE with the IDENTIFIER "y". The variable should be of the UNDERLYING of your custom TYPE "x"

// 2. In func ex05

// 🅐 this should already be done
//	 ① print out the value of the variable "x"
//	 ⓶ print out the type of the variable "x"
//	 ⓷ assign your own VALUE to the VARIABLE "x" using the "=" OPERATOR
// 	 ⓸ print out the value of the variable "x"

// 🅑 now do this
//	 ① now use CONVERSION to convert the TYPE of the VALUE stored in "x" to the UNDERLYING TYPE
//	 ⓶ then use the short declaration operator to ASSIGN that value to "y"
//	 ⓷ print out the value stored in "y"
// 	 ⓸ print out the type of "y"

package main

import "fmt"

type numbers int
var x numbers
var y int

func ex05() {
	fmt.Println(x) // 0
	fmt.Printf("%T\n", x) // main.number

	x = 90
	fmt.Println(x) // 90

	y = int(x)
	fmt.Println(y) // 90
	fmt.Printf("%T\n", y) // int
	fmt.Printf("%T\n", x) // main.number

}