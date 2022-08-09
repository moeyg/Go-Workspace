// numeric : 숫자 타입

package main

import "fmt"

func numeric() {
	// uint : 부호 없는 정수, 32bit 시스템에서는 uint32 / 64bit 시스템에서는 uint64
	// uint8 : 0 ~ 255
	// uint16 : 0 ~ 65535
	// uint32 : 0 ~ 4294967295
	// uint64 : 0 ~ 18446744073709551615

	var x byte // byte = uint8
	fmt.Printf("%T\n", x) // uint8

	// 2진수 정수 표현 = 부호 비트(0/1) + 15자리
	// int : 부호 있는 정수, 32bit 시스템에서는 uint32 / 64bit 시스템에서는 uint64
	// int8 : -128 ~ 127
	// int16 : -32768 to 32767
	// int32 : -2147483648 ~ 2147483647
	// int64 : -9223372036854775808 ~ 9223372036854775807

	var y rune // rune = int32 -> 문자가 4바이트를 사용하기 때문
	fmt.Printf("%T\n", y) // int32

	// 2진수 실수 표현 = 부호 비트(0/1) + 지수부(8비트) + 소수부(23비트)
	// float32 : 부호 있는 실수로 소수부 7자리 (1.175494351 * 10^-38 ~ 3.402823466 * 10^38)
	// float64 : 부호 있는 실수로 소수부 15자리 (2.2250738585072014 * 10^-308 ~ 1.7976931348634158 * 10^308)

	// complex64 : 8바이트 복소수
	// complex128

}