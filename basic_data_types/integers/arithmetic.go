package main

import "fmt"

func main() {
	// If the result of an arithmetic operation, whether
	// signed or unsigned, has more bits than can be represented in
	// result type, it is said to overflow.
	// The high-order bits that do not fit are silently discarded.
	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // "255 0 1"

	// If the original number is a signed type, the result could be
	// negative if the leftmost bit is a 1
	var i int8 = 127
	fmt.Println(i, i+1, i*i) // "127 -128 1"
}
