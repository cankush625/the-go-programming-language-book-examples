package main

import "fmt"

func eliminateAdjacentDuplicates(strings []string) []string {
	idx := 0 // Index of last written string
	for _, element := range strings {
		if strings[idx] == element {
			continue
		}
		idx++
		strings[idx] = element
	}
	return strings[:idx+1]
}

func main() {
	data := []string{"first", "", "second", "second", "", "", "third"}
	result := eliminateAdjacentDuplicates(data)
	fmt.Printf("%q\n", data)   // "["first" "" "second" "" "third" "" "third"]"
	fmt.Printf("%q\n", result) // "["first" "" "second" "" "third"]"
}
