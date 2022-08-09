// string : 문자 타입

package main

import "fmt"

func string() {
	// string은 byte로 이루어진 문자열
	str := "Hello World"
	fmt.Println(str) // Hello World
	fmt.Printf("%T\n", str) // string

	// slice byte
	sliceStr := []byte(str)
	fmt.Println(sliceStr) // [72 101 108 108 111 32 87 111 114 108 100]
	fmt.Printf("%T\n", sliceStr) // []uint8
}