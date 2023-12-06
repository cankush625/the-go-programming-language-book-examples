package main

import "fmt"

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for key, xVal := range x {
		if yVal, ok := y[key]; !ok || xVal != yVal {
			return false
		}
	}
	return true
}

func main() {
	map1 := map[string]int{"A": 10}
	map2 := map[string]int{"A": 10}
	map3 := map[string]int{"B": 10}
	fmt.Printf("%t\n", equal(map1, map2)) // "true"
	fmt.Printf("%t\n", equal(map1, map3)) // "false"
}
