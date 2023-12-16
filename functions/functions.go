package main

import "fmt"

// Function can be defined using parameters and return type
func add(num1 int, num2 int) int {
	return num1 + num2
}

// Function can be defined using named return type
// Each name if the result/return declares a local variable
// initialized to the zero value for it's type
func subtract(num1, num2 int) (res int) {
	res = num1 - num2
	return
}

// The black identifier can be used to indicate that a
// parameter is unused
func first(num1 int, _ int) int {
	return num1
}

func zero(int, int) int {
	return 0
}

// Note:
// The type of function is sometimes called its signature.
// Two functions have the same type or signature if they have
// the same sequence of parameter types and the same sequence
// of result types. The names of parameters and results don't
// affect the type, nor does whether they were declared using
// the factored form.

// We may encounter a function declaration without a body,
// indicating that the function is implemented in a language
// other than Go. Such a declaration defines the function
// signature.

func main() {
	fmt.Printf("%T\n", add)      // "func(int, int) int"
	fmt.Printf("%T\n", subtract) // "func(int, int) int"
	fmt.Printf("%T\n", first)    // "func(int, int) int"
	fmt.Printf("%T\n", zero)     // "func(int, int) int"
}
