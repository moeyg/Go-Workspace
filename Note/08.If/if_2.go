package main

import "fmt"

func if_2() {
	month := 5

	if 1 <= month && month <= 12 { 
		if 3 <= month && month <= 5 {
			fmt.Println("Spring")
		} else if 6 <= month && month <= 8 {
			fmt.Println("Summer")
		} else if 9 <= month && month <= 11 {
			fmt.Println("Autumn")
		} else {
			fmt.Println("Winter")
		}
	} else {
		fmt.Println("A year is from 1 to 12.")
	}
}