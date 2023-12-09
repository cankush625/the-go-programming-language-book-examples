package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

// Count number of UTF-8 characters
func main() {
	counts := make(map[string]map[rune]int) // counts of unicode characters
	var utflen [utf8.UTFMax + 1]int         // count of lengths of UTF-8 encodings
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune. nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			letters := counts["letter"]
			if letters == nil {
				letters := make(map[rune]int)
				letters[r]++
				counts["letter"] = letters
			} else {
				letters[r]++
			}
		} else if unicode.IsDigit(r) {
			digits := counts["digit"]
			if digits == nil {
				digits := make(map[rune]int)
				digits[r]++
				counts["digit"] = digits
			} else {
				digits[r]++
			}
		} else {
			others := counts["other"]
			if others == nil {
				others := make(map[rune]int)
				others[r]++
				counts["other"] = others
			} else {
				others[r]++
			}
		}
		utflen[n]++
	}
	for t, data := range counts {
		fmt.Printf("type %s\n", t)
		fmt.Printf("rune\tcount\n")
		for c, n := range data {
			fmt.Printf("%q\t%d\n", c, n)
		}
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

// Note:
// Optimized implementation: https://github.com/torbiak/gopl/blob/master/ex4.8/charcount.go

// Input:
// Hi, I'm Ankush Chavan. 1145
// Output:
// type digit
// rune    count
// '1'     2
// '4'     1
// '5'     1
// type letter
// rune    count
// 'h'     2
// 'C'     1
// 'a'     2
// 'v'     1
// 'A'     1
// 'H'     1
// 'm'     1
// 'n'     2
// 'k'     1
// 'i'     1
// 'I'     1
// 'u'     1
// 's'     1
// type other
// rune    count
// ','     1
// ' '     4
// '\''    1
// '.'     1
// '\n'    1
//
// len     count
// 1       28
// 2       0
// 3       0
// 4       0
