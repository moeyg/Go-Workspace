// 1. Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS "integer" and "str" and "boolean"

// 🅐 90
// 🅑 "Go"
// 🅒 true

// 2. Now print the value stored in those variable using

// 🅐 a single print statement
// 🅑 multiple print statement

package main

import "fmt"

func ex01() {
	integer := 90
	str := "Go"
	boolean := true

	fmt.Println(integer) // 90
	fmt.Println(str) // Go
	fmt.Println(boolean) // true

	fmt.Println(integer, str, boolean) // 90 Go true
}