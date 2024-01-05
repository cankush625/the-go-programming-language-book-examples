package main

import "fmt"

func sqlQuote(x interface{}) string {
	if x == nil {
		return "NULL"
	} else if _, ok := x.(int); ok {
		return fmt.Sprintf("%d", x)
	} else if _, ok := x.(uint); ok {
		return fmt.Sprintf("%d", x)
	} else if b, ok := x.(bool); ok {
		if b {
			return "TRUE"
		}
		return "FALSE"
	} else if s, ok := x.(string); ok {
		return s
	} else {
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}

func main() {
	fmt.Println(sqlQuote(10))   // "10"
	fmt.Println(sqlQuote(true)) // "TRUE"
}
