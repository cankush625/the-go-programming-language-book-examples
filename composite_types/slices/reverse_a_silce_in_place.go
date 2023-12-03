package main

import "fmt"

func reverse(s []int) {
	for start, end := 0, len(s)-1; start < end; start, end = start+1, end-1 {
		s[start], s[end] = s[end], s[start]
	}
}

func main() {
	numbers := []int{2, 5, 3, 7, 8}
	fmt.Println(numbers) // "[2 5 3 7 8]"
	reverse(numbers)
	fmt.Println(numbers) // "[8 7 3 5 2]"
}
