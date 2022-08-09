// Bit Shift Operator (비트 시프트 연산자)

package main

import "fmt"

func bitShift() {
	num := 4
	fmt.Printf("%d binary : %b\n", num, num) // 4 binary : 100

	// Left Shift : <<
	leftShift := num << 1
	fmt.Printf("%d binary : %b\n", leftShift, leftShift) // 8 binary : 1000

	// Right Shift : >>
	rightShift := num >> 1
	fmt.Printf("%d binary : %b\n", rightShift, rightShift) // 2 binary : 10

	const (
		_ = iota
		kb = 1 << (iota * 10)
		mb = 1 << (iota * 10)
		gb = 1 << (iota * 10)
	)
	fmt.Printf("%d binary : %b\n", kb, kb) // 1024 binary : 10000000000
	fmt.Printf("%d binary : %b\n", mb, mb) // 1048576 binary : 100000000000000000000
	fmt.Printf("%d binary : %b\n", gb, gb) // 1073741824 binary : 1000000000000000000000000000000
}