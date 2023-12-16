package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	freq := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		freq[word]++
	}
	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, scanner.Err())
		os.Exit(1)
	}
	for word, n := range freq {
		fmt.Printf("%-30s %d\n", word, n)
	}
}

// Input:
// Go is a good programming language. Go is faster and intuitive than most other programming languages.
// Output:
// good                           1
// language.                      1
// most                           1
// programming                    2
// faster                         1
// than                           1
// other                          1
// languages.                     1
// a                              1
// and                            1
// Go                             2
// is                             2
// intuitive                      1
