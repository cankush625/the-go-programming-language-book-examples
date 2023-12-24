package main

import (
	"fmt"
)

func foo() {
	defer fmt.Println("Bye foo")
	fmt.Println("Hello foo")
	bar()
	fmt.Println("After bar")
}

func bar() {
	defer fmt.Println("Bye bar")
	fmt.Println("Hello bar")
	baz()
	fmt.Println("After baz")
}

func baz() {
	defer recoverPanic()
	defer fmt.Println("Bye baz")
	fmt.Println("Hello baz")
	panic("panic in baz")
}

func recoverPanic() {
	if err := recover(); err != nil {
		fmt.Printf("RECOVERED: %v\n", err)
	}
}

func main() {
	foo()
	fmt.Println("After foo")
	// Note:
	// The recover only works in the deferred function. This is
	// because during the returning to the caller in panicking
	// sequence, the only code that runs is deferred functions,
	// so the recover will not run elsewhere.
	// Ref: https://gosamples.dev/recover/#
}

// Output:
// Hello foo
// Hello bar
// Hello baz
// Bye baz
// RECOVERED: panic in baz
// After baz
// Bye bar
// After bar
// Bye foo
// After foo
