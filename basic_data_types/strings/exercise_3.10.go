package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Using comma_iterative
	fmt.Println(comma_iterative("123"))     // "123"
	fmt.Println(comma_iterative("12345"))   // "12,345"
	fmt.Println(comma_iterative("1234567")) // "1,234,567"
}

func comma_iterative(number string) string {
	var buf bytes.Buffer
	start_digits := len(number) % 3
	if start_digits == 0 && len(number) >= 3 {
		start_digits = 3
	}
	// Write the first group of numbers upto 3 digits
	buf.WriteString(number[:start_digits])
	// Remaining numbers will exactly be in a bunch of 3
	for i := start_digits; i < len(number); i += 3 {
		buf.WriteString("," + number[i:i+3])
	}
	return buf.String()
}
