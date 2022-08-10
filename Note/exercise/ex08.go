// Create TYPED and UNTYPED constants. Print the values of the constants.

package main

import "fmt"

const (
	// TYPED constant
	numTyped = 90
	// UNTYPED constant
	numUntyped int = 100
)

func ex08() {
	fmt.Println(numTyped) // 90
	fmt.Println(numUntyped) // 100
}