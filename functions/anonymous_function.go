package main

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	// Function values are not just code but can have state.
	// The anonymous inner function in squares func can access
	// and update the local variables of the enclosing function squares.

	// Function values like these are implemented using a technique
	// called closures, and Go programmers often use this term for
	// function values.
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
}
