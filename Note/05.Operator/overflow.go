// Overflow : 정수의 가장 큰 값에 값을 더하면 가장 작은 값부터 시작된다.

package main

import "fmt"

func overflow() {
	var num int8 = 127

	fmt.Printf("%08b(%d)\n", uint8(num + 1), num + 1) // 10000000(-128)
	fmt.Printf("%08b(%d)\n", uint8(num + 2), num + 2) // 10000001(-127)
	fmt.Printf("%08b(%d)\n", uint8(num + 3), num + 3) // 10000010(-126)
	fmt.Printf("%08b(%d)\n", uint8(num + 10), num + 10) // 10001001(-119)
}