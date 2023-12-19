package main

import "fmt"

func square(n int) int {
	return n * n
}

// Functions can be passed as function parameters
func doubleSquare(n int, square func(int) int) int {
	if square != nil {
		return square(n) * 2
	}
	return 0
}

func main() {
	// Functions are first-class values in Go,
	// like other values, function values have types, and
	// they may be assigned to variables or passed to or
	// returned from functions.
	s := square
	fmt.Println(s(10)) // "100"
	// Zero values of a function type is nil. Calling a nil
	// function value causes a panic.
	// Function values may be compared with nil; but they are not
	// comparable, so they may not be compared against each other
	// or used as keys in a map
	var f func(int) int
	if f != nil {
		f(5)
	}
	// Function can be passed as a argument to the function
	fmt.Println(doubleSquare(10, square)) // "200"
}
