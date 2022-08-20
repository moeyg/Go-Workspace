package main

import "fmt"

func switchBoolean() {
	temperature := 25

	// true 값을 넣어 boolean 값을 검사할 수 있다.
	switch true {
		case temperature > 27 :
			fmt.Println("Summer")
		case temperature < 3 :
			fmt.Println("Winter")
		case temperature > 12 :
			fmt.Println("Spring")
		default :
		fmt.Println("Autumn")
	}	
}