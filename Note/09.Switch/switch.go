// Switch : 값을 검사하는 조건문

package main

import "fmt"

func switchStatement() {
	month := 1
	
	// switch 값 { }
	switch month {
		// ,를 사용해 값을 여러개 비교할 수 있다.
		case 1, 2, 12 :
			fmt.Println("winter")
		case 3, 4, 5 :
			fmt.Println("Spring")
		case 6, 7, 8 :
			fmt.Println("Summer") 
		case 9, 10, 11 :
			fmt.Println("Autumn")
		default :
			fmt.Println("A year is from 1 to 12.")
	}
}