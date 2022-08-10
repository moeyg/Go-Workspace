// Create a variable of type string using a raw string literal. Print it.

package main

import "fmt"

func ex10() {
	golang := `Go is a statically typed, compiled programming language designed at "Google" by Robert Griesemer, Rob Pike, and Ken Thompson.
	It is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency.`

	fmt.Println(golang)
	/*
	Go is a statically typed, compiled programming language designed at "Google" by Robert Griesemer, Rob Pike, and Ken Thompson.
        It is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency.
	*/
}