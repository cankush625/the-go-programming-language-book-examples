package main

import "fmt"

func makeKey(list []string) string {
	return fmt.Sprintf("%q", list)
}

func countSlices(list [][]string) map[string]int {
	var m = make(map[string]int)
	for _, lst := range list {
		m[makeKey(lst)]++
	}
	return m
}

func main() {
	s := [][]string{[]string{"a", "b", "c"}, []string{"a", "g", "c"}, []string{"a", "b", "c"}}
	result := countSlices(s)
	for k, v := range result {
		fmt.Printf("%s is present %d times.\n", k, v)
	}
}

// Output:
// ["a" "b" "c"] is present 2 times.
// ["a" "g" "c"] is present 1 times.
