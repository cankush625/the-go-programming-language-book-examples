package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	// If all the fields of the struct are comparable then the
	// struct is also comparable
	p := Point{1, 2}
	q := Point{1, 2}
	r := Point{2, 3}

	fmt.Println(p == q) // "true"
	fmt.Println(p == r) // "false"

	// Comparable struct types, like other comparable types, may be
	// used as the key type of a map
	occurrence := make(map[Point]int)
	occurrence[p]++
	occurrence[q]++
	occurrence[r]++
	for key, val := range occurrence {
		fmt.Printf("Point {%d, %d}\t%d\n", key.X, key.Y, val)
	}
	// Output:
	// Point {1, 2}    2
	// Point {2, 3}    1
}
