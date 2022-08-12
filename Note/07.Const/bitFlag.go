// Bit Flag : 비트 플래그, 각 비트마다 의미를 부여하는 것

package main

import "fmt"

const (
	BitFlag1 uint = 1 << iota // 1 = 1 << 0
	BitFlag2				  // 2 = 1 << 1
	BitFlag3				  // 4 = 1 << 2
	BitFlag4				  // 8 = 1 << 3
)

func bitFlag() {
	fmt.Printf("%d : %08b\n", BitFlag1, BitFlag1) // 1 : 00000001
	fmt.Printf("%d : %08b\n", BitFlag2, BitFlag2) // 2 : 00000010
	fmt.Printf("%d : %08b\n", BitFlag3, BitFlag3) // 4 : 00000100
	fmt.Printf("%d : %08b\n", BitFlag4, BitFlag4) // 8 : 00001000
}