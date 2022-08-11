// Floating point error : 실수 오차
// 실수인 값을 더하면 값이 오차가 나서 정확한 계산 X
// 방법 1. Nextafter() 사용

package main

import (
	"fmt"
	"math"
)

func equal(a, b float64) bool {
	return math.Nextafter(a, b) == b
}

func pointError() {
	var x float64 = 0.1
	var y float64 = 0.2
	var z float64 = 0.3

	fmt.Printf("%f + %f == %f : %v\n", x, y, z, x + y == z) // 0.100000 + 0.200000 == 0.300000 : false
	fmt.Println(x + y) // 0.30000000000000004

	// Nextafter()
	fmt.Printf("%0.18f == %0.18f : %v\n", z, x + y, equal(x + y, z)) 
	// 0.299999999999999989 == 0.300000000000000044 : true
}