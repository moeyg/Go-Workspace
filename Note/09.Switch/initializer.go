// Initializer : 초기문

package main

import "fmt"

func IsAge() (age int) {
	fmt.Println("How old are you? :")
	fmt.Scanf("%d", &age)
	return age
}

// main 함수
func initializer() {
	// switch 초기문; 비교값 { }
	// 변수 age는 switch 내에서 사용하는 지역 변수 
	switch age := IsAge(); true {
		case age < 0 :
			fmt.Println("Please enter your correct age")

		case 0 <= age && age < 5 :
			fmt.Printf("%d years old - baby\n", age)
			fmt.Println("FREE ENTRY")

		case age < 10 :
			fmt.Printf("%d years old - child\n", age)
			fmt.Println("PAY $35")

		case age < 20 :
			fmt.Printf("%d years old - teenager\n", age)
			fmt.Println("PAY $50")

		case age < 60 :
			fmt.Printf("%d years old - adult\n", age)
			fmt.Println("PAY $60")
		
		case age >= 60:
			fmt.Printf("%d years old, senior\n", age)
			fmt.Println("FREE ENTRY")
		
		default :
			fmt.Println("Please enter your correct age")
	}
}