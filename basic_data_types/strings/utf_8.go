package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str_chinese := "你好, world"
	fmt.Println(len(str_chinese)) // "13"

	// Get rune count in string
	fmt.Println(utf8.RuneCountInString(str_chinese)) // "9"

	// Decode rune in string
	for i := 0; i < len(str_chinese); {
		r, size := utf8.DecodeRuneInString(str_chinese[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	// Output:
	// 0       你
	// 3       好
	// 6       ,
	// 7
	// 8       w
	// 9       o
	// 10      r
	// 11      l
	// 12      d

	// Converting an integer value to a string interprets
	// the integer as a rune value and yield the UTF-8
	// representation of that rune
	fmt.Println(string(65)) // "A"

	// If rune is invalid, the replacement character is substituted
	fmt.Println(string(1234567)) // "�"
}
