// Type Range : 타입 범위

// uint8 : 1byte(8bit) 부호 없는 정수 (0 ~ 255)
// uint16 : 2byte(16bit) 부호 없는 정수 (0 ~ 65535)
// uint32 : 4byte(32bit) 부호 없는 정수 (0 ~ 4294967295)
// uint64 : 8byte(64bit) 부호 없는 정수 (0 ~ 18446744073709551615)

// 2진수 정수 표현 = 부호 비트(0/1) + 15자리
// int8 : 1byte(8bit) 부호 있는 정수 (-128 ~ 127)
// int16 : 2byte(16bit) 부호 있는 정수 (-32768 to 32767)
// int32 : 4byte(32bit) 부호 있는 정수 (-2147483648 ~ 2147483647)
// int64 : 8byte(64bit) 부호 있는 정수 (-9223372036854775808 ~ 9223372036854775807

// 2진수 실수 표현 = 부호 비트(0/1) + 지수부(8비트) + 소수부(23비트)
// float32 : 4byte(32bit) 부호 있는 실수로 소수부 7자리 (1.175494351 * 10^-38 ~ 3.402823466 * 10^38)
// float64 : 8byte(64bit) 부호 있는 실수로 소수부 15자리 (2.2250738585072014 * 10^-308 ~ 1.7976931348634158 * 10^308)

// complex64 : 복소수
// complex128

// byte : uint8
// rune : int32

package main

import "fmt"

func typeRange() {
	// 큰 타입에서 작은 타입으로 변환할 때 값이 삭제될 수 있다.
	var integer16 = 3456
	var integer8 = int8(integer16)
	fmt.Println(integer16) // 3456 -> 00001101 10000000
	fmt.Println(integer8) // -128 -> (00001101 삭제) 10000000

	var float1 float32 = 1234.567
	var float2 float32 = 3456.789
	var float3 = float1 * float2
	var float4 = float3 * 3

	// 자릿수 오차가 발생할 수 있기 때문에 계산에 유의한다.
	fmt.Println(float3) // 4.267638e+06 = 4267637.625363
	fmt.Println(float4) // 1.2802914e+07 = 12802912.876089
}