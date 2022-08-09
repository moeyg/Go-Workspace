// Type Range : 타입 범위

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