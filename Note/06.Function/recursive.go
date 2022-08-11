// Recursive : 재귀함수
// 재귀 호출을 할 때마다 stack이 쌓임
// Go는 자동 증가 stack을 사용 -> 메모리가 고갈되면 자동 종료

package main

import "fmt"

func PrintNum (n int) {
	// 재귀함수 탈출 조건 -> 보장이 되지 않으면 무한 반복
	if n == 0 {
		return
	}

	fmt.Println(n)
	// 재귀
	PrintNum(n - 1)
	// 재귀함수를 호출했던 값으로 돌아감
	fmt.Println("Number : ", n)
}

func main() {
	PrintNum(3)
}

// 18번 줄에서 PrintNum(3) 호출 -> 12번 줄에서 3 출력
// 3

// 13번 줄에서 PrintNum(3 - 1) 호출 -> 12번 줄에서 2 출력
// 2

// 13번 줄에서 PrintNum(2 - 1) 호출 -> 12번 줄에서 1 출력
// 1

// 8번 줄에서 n == 0, 9번 줄에서 return -> 호출된 13번 줄로 이동
// 호출한 13번째 줄로 돌아가서 PrintNum(1) -> 14번 줄 출력
// Number :  1

// (함수 끝) -> PrintNum(1)을 호출한 13번째 줄로 돌아가서 PrintNum(2) -> 14번 줄 출력)
// Number :  2

// (함수 끝) -> PrintNum(2)를 호출한 13번째 줄로 돌아가서 PrintNum(3) -> 14번 줄 출력)
// Number :  3