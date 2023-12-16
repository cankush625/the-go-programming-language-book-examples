package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main() {
	// Make a slice of non-empty string elements
	data := []string{"first", "", "third"}
	// Print the result of nonempty function
	fmt.Printf("%q\n", nonempty(data)) // "["first" "third"]
	// Print the original data slice
	fmt.Printf("%q\n", data) // ["first" "third" "third"]
}
