package main

import "fmt"

func main() {
	// left shift operator
	var left_shift_once uint8 = 1 << 1
	var left_shift_five_times uint8 = 1 << 5
	// OR operator
	var x = left_shift_once | left_shift_five_times
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x) // "00100010"
	fmt.Printf("%08b\n", y) // "00000110"

	// Bitwise AND
	fmt.Printf("%08b\n", x&y) // "00000010"

	// Bitwise OR
	fmt.Printf("%08b\n", x|y) // "00100110"

	// Bitwise XOR
	fmt.Printf("%08b\n", x^y) // "00100100"

	// Bit clear operator (AND NOT)
	fmt.Printf("%08b\n", x&^y) // "00100000"
}
