// Special Symbols (특수문자)

package main

import "fmt"

func specialSymbols() {
	var integer int = 90

	// \n : 개행(줄바꿈)
	fmt.Printf("\\n : %d\n", integer) // \n : 90

	// \t : 탭
	fmt.Printf("\\t : \t%d\n", integer) // \t :    90

	// \\ : \ 출력
	fmt.Printf("\\\\ : \\%d\\\n", integer) // \\ : \90\

	// \' \" : ' 과 " 출력
	fmt.Printf("\\\" : \"%d\"\n", integer) // \" : "90"
}