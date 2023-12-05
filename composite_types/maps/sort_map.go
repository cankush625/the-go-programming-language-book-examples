package main

import (
	"fmt"
	"sort"
)

func sortMap(data map[string]int) {
	// Sorting map requires first sorting the
	// keys of the map and then iterating over these
	// sorted keys to get the map element in sorted order
	var keys []string
	for key := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("%s\t%d\n", key, data[key])
	}
}

func main() {
	data := map[string]int{
		"bob":   10,
		"alice": 15,
	}
	sortMap(data)
}
