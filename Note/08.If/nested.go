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

func isValidUserData() bool {
	if id == userID {
		if pw == userPW {
			fmt.Println("안녕하세요!")
			return true
		}
		fmt.Println("비밀번호가 틀렸습니다.")
		return false
	}
	if id != userID {
		fmt.Println("아이디가 일치하지 않습니다.")
	}
	return false
}

func main() {
	fmt.Println("Please enter your ID :")
	fmt.Scanln(&id)
	fmt.Println("Please enter your PW :")
	fmt.Scanln(&pw)
	
	isValidUserData()
}