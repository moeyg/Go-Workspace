// Conditional Statement : 조건문

package main

import "fmt"

func if_1() {
	temperature := 2

	if temperature > 27 {
		fmt.Println("Summer")
	} else if temperature < 3 {
		fmt.Println("Winter")
	} else if temperature > 12 {
		fmt.Println("Spring")
	} else {
		fmt.Println("Autumn")
	}
}