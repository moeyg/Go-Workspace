// iota : const 키워드와 괄호 안에서 iota를 이용해 상수를 열거할 수 있다.
// 상수명의 Zero Value에 iota를 지정하면 첫번째 상수값은 0으로 정의하고 이후의 상수값은 순차적으로 1씩 증가한 값을 갖는다.

package main

import "fmt"

const (
	zero = iota
	one
	two
	three
)

const (
	JAN = iota + 1
	FEB
	MAR
)

func iotaDeclaration() {
	fmt.Println(zero) // 0
	fmt.Println(one) // 1
	fmt.Println(two) // 2
	fmt.Println(three) // 3

	fmt.Printf("%d월\n", JAN) // 1월
	fmt.Printf("%d월\n", FEB) // 2월
	fmt.Printf("%d월\n", MAR) // 3월
}