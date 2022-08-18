// Nested if : 중첩 if문
// 음식값이 오만 원이 넘고 친구 중에

package main

import "fmt"

const (
	userID = "Gopher"
	userPW = "90lang"
)

var id string
var pw string

func IsValid(boolean bool) {
	if boolean == true {
		fmt.Println("안녕하세요!")
	}
	if boolean == false {
	fmt.Println("ID 또는 비밀번호가 일치하지 않습니다.")
	}
	return
}

func IsValidUserData() bool {
	if id != userID {
		return false
	}
	if id == userID {
		if pw == userPW {
			return true
		}
	}
	return false
}

// main 함수
func nested() {
	fmt.Println("Please enter your ID :")
	fmt.Scanln(&id)
	fmt.Println("Please enter your PW :")
	fmt.Scanln(&pw)
	
	var isValid bool = IsValidUserData()
	IsValid(isValid)
}