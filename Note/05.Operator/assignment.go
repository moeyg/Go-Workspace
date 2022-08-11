// Assignment Operator : 대입 연산자
// Go는 값을 반환 X -> 변수에 저장 X = 전위증감, 후위증감 X

package main

import (
	"fmt"
)

func main() {
	var x int
	var y int

	// 값 대입
	x, y = 10, 20
	fmt.Printf("x : %d, y : %d\n", x, y) // x : 10, y : 20

	// 값 교환
	x, y = y, x
	fmt.Printf("x : %d, y : %d\n", x, y) // x : 20, y : 10

	// 복합 대입 연산자
	x += 10
	fmt.Println(x) // 30

	x -= 10
	fmt.Println(x) // 20

	x *= 10
	fmt.Println(x) // 200

	x /= 10
	fmt.Println(x) // 20

	y %= 3
	fmt.Println(y) // 1

	// 증감 연산자
	a := 10
	a ++
	fmt.Println(a) // 11
	a --
	fmt.Println(a) // 10

	// 연산자 우선 순위
	// 1. *, /, %, <<, >>, &, &^
	// 2. +, -, |, ^
	// 3. ==, !=, <, <=, >, >=
	// 4. &&
	// 5. ||
	fmt.Println(3 * 4 ^ 7 << 2 + 3 * 5 == 7) // false
	// (3 * 4) ^ (7 << 2) + (3 * 5) == 7
	// (12 ^ 28) + 15 == 7
	// (0000 1100 ^ 0001 1100) + 15 == 7
	// 0001 0000(16) + 15 == 7
	// 31 == 7 -> false
}
