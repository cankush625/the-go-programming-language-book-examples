package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// Using basename
	fmt.Println(basename("a/b/c.go")) // "c"
	fmt.Println(basename("c.d.go"))   // "c.d"
	fmt.Println(basename("abc"))      // "abc"

	// Using comma
	fmt.Println(comma("123"))     // "123"
	fmt.Println(comma("12345"))   // "12,345"
	fmt.Println(comma("1234567")) // "1,234,567"

	// Using intsToString
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
}

func basename(filename string) string {
	slash := strings.LastIndex(filename, "/") // -1 if "/" not found
	filename = filename[slash+1:]
	if dot := strings.LastIndex(filename, "."); dot >= 0 {
		filename = filename[:dot]
	}
	return filename
}

func comma(number string) string {
	n := len(number)
	if n <= 3 {
		return number
	}
	return comma(number[:n-3]) + "," + number[n-3:]
}

func intsToString(numbers []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range numbers {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
