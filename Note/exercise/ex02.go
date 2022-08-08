// 1. Use var to DECLARE three variables. The variables should have package level scope.
// Do not assign VALUE to the variables. Use the following IDENTIFIERS for the variables and make sure the variables are store value of the following TYPE.

// 🅐 identifier "initInt" type int
// 🅑 identifier "initString" type string
// 🅒 identifier "initBool" type bool

// 2. In func ex02

// 🅐 print out the values for each identifier
// 🅑 The compiler assigned values to the variables. What are these values called?

package main

import "fmt"

var initInt int
var initString string
var initBool bool

func ex02() {
	fmt.Println(initInt) // 0
	fmt.Println(initString) // ""
	fmt.Println(initBool) // false

	fmt.Printf("integer type is %T\n", initInt) // iinitInt type is int
	fmt.Printf("str type is %T\n", initString) // initString type is string
	fmt.Printf("boolean type is %T\n", initBool) // initBool type is bool
}