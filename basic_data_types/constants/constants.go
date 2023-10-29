package main

import "fmt"

func main() {
	const pi = 3.142
	fmt.Println(pi) // 3.142

	// When sequence of constants is declared as a group,
	// the right-hand side expression may be omitted for all
	// but the first of the group, implying that the previous
	// expression, and it's type should be used again
	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d) // "1 1 2 2"
}
