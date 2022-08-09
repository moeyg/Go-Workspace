package main

import "fmt"

func iotaFunc() {
	// iota : 데이터 타입이 없는 정수를 연속으로 생성
	const (
		a = iota
		b = iota
		c
		d
	)
	fmt.Println(a) // 0
	fmt.Println(b) // 1
	fmt.Println(c) // 2
	fmt.Println(d) // 3

	const (
		x = iota + 10
		y
		z
	)
	fmt.Println(x) // 10
	fmt.Println(y) // 11
	fmt.Println(z) // 12

	const (
		Jan = iota + 1
		Feb
		Mar
		Apr
		May
		Jun
		Jul
		Aug
		Sep
		Oct
		Nov
		Dec
	)
	fmt.Println(Jan) // 1
	fmt.Println(Feb) // 2
	fmt.Println(Mar) // 3
	fmt.Println(Apr) // 4
	fmt.Println(May) // 5
	fmt.Println(Jun) // 6
	fmt.Println(Jul) // 7
	fmt.Println(Aug) // 8
	fmt.Println(Sep) // 9
	fmt.Println(Oct) // 10
	fmt.Println(Nov) // 11
	fmt.Println(Dec) // 12
}