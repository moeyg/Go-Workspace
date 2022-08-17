// Float Output (실수 출력) : %f, %g
// fmt.Printf("%f", float)
// fmt.Printf("%g", float)
// fmt.Printf("%v", float)

package main

import "fmt"

func float() {
	var pi float64 = 3141592.65359
	
	// %f : 소수점 6자리까지 표현
	fmt.Printf("f : %f\n", pi) // f : 3141592.653590

	// %g(%v) : 길이(6자리)에 따라서 지수 표현 또는 실수 표현으로 출력
	fmt.Printf("g : %g\n", pi) // g : 3.14159265359e+06
	fmt.Printf("v : %v\n", pi) // v : 3.14159265359e+06
}