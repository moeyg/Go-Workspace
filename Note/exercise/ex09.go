// Write a program that

// 🅐 assign an int to a variable
// 🅑 prints that int in decimal, binary, and hex
// 🅒 shifts the bits of that int over 1 position to the left, and assigns that to variable
// 🅓 prints that variable in decimal, binary, and hex

package main

import "fmt"

func ex09() {
	num := 90
	fmt.Printf("%d\n", num) // 90
	fmt.Printf("%b\n", num) // 1011010
	fmt.Printf("%#x\n", num) // 0x5a

	numShift := num << 1
	fmt.Printf("%d\n", numShift) // 180
	fmt.Printf("%b\n", numShift) // 10110100
	fmt.Printf("%#x\n", numShift) // 0xb4
}