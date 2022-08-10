// Using iota, create 4 constants for the 4 years. Print the constant values

package main

import "fmt"

const (
	year2022 = 2022 + iota
	year2023 = 2022 + iota
	year2024 = 2022 + iota
	year2025 = 2022 + iota
)

func main() {
	fmt.Println(year2022) // 2022
	fmt.Println(year2023) // 2023
	fmt.Println(year2024) // 2024
	fmt.Println(year2025) // 2025
}