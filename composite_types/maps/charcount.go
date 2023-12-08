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
	counts := make(map[rune]int)    // counts of unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
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
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
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

// Input:
// Hi, I'm Ankush Chavan.
// Output:
// rune    count
// 's'     1
// 'C'     1
// '.'     1
// ' '     3
// 'n'     2
// 'm'     1
// 'A'     1
// 'u'     1
// 'i'     1
// 'I'     1
// '\n'    2
// ','     1
// 'a'     2
// 'k'     1
// 'h'     2
// 'v'     1
// 'H'     1
// '\''    1
//
// len     count
// 1       24
// 2       0
// 3       0
// 4       0
