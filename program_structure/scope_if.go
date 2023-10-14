package main

import "fmt"

func main() {
	if x := f(); x == 0 {
		fmt.Println(x)
	} else if y := g(x); x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}
	// We cannot use fmt.Println(x, y) here
	// because x and y are only access in the
	// scope of if statement.
}

func f() int {
	return 10
}

func g(num int) int {
	return num
}
