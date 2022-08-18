// Initializer : 초기화
// if 초기문; 조건문 { }
// 초기문 -> 조건문

package main

import "fmt"

func UploadFile() (filename string, success bool) {
	return filename, success
}

// main 함수
func initializer() {
	if filename, success := UploadFile(); success {
		fmt.Println("Upload success", filename)
	} else {
		fmt.Println("Failed to upload")
	}
}