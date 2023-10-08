package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLinesWithFile(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				continue
			}
			countLinesWithFile(f, counts)
			f.Close()
		}
	}
	for filename, data := range counts {
		for line, count := range data {
			if count > 1 {
				fmt.Printf("%d\t%s\t%s\n", count, line, filename)
			}
		}
	}
}

func countLinesWithFile(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[f.Name()] == nil {
			counts[f.Name()] = make(map[string]int)
		}
		counts[f.Name()][input.Text()]++
	}
}
