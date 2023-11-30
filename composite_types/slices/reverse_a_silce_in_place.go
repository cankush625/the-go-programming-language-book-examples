package main

import "fmt"

func reverse(s []int) {
	for start, end := 0, len(s)-1; start < end; start, end = start+1, end-1 {
		s[start], s[end] = s[end], s[start]
	}
}

func main() {
	numbers := []int{2, 5, 3, 7, 8}
	fmt.Println(numbers)
	reverse(numbers)
	fmt.Println(numbers)
}
