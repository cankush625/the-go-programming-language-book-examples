package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "hello, world"

	// Built-in function len returns number of bytes
	// (not runes) in a string
	fmt.Println(len(str)) // "12"

	// Index operation str[i] returns i-th byte of string
	fmt.Println(str[3]) // "108"  ('l')

	// Getting substring
	fmt.Println(str[7:]) // "world"

	// If we have string whose single character requires
	// more bytes to represent than english characters
	// then we can see the difference clearly that one
	// chinese character requires 3 bytes, so using len
	// to get length of string is not useful
	// Here, we need to count the runes in the string. There
	// are only nine code points or runes in below string
	str_chinese := "你好, world"
	fmt.Println(len(str_chinese))                    // "13"
	fmt.Println(utf8.RuneCountInString(str_chinese)) // "9"
}
