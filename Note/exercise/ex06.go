// Write a program that prints a number in decimal, binary, and hex

package main

import "fmt"

func ex06() {
	num := 90

	// decimal
	fmt.Printf("%d\n", num) // 90
	// binary
	fmt.Printf("%b\n", num) // 1011010
	// hex
	fmt.Printf("%#x\n", num) // 0x5a
}