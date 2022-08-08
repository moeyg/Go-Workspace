// 1. At the package level scope, assign the following values to the three variables

// 🅐 for integer assign 90
// 🅑 for str assign "Go"
// 🅒 for boolean assign true

// 2. In func ex03

// 🅐 use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDNETIFIER "variables"
// 🅑 print out the value stored by variable "variables"

package main

import "fmt"

var integer int = 90
var str string = "Go"
var boolean bool = true

func ex03() {
	variables := fmt.Sprintf("%v\t%v\t%v", integer, str, boolean)
	fmt.Println(variables)
}