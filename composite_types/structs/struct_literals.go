package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

func ScalePoint(point Point, factor int) Point {
	return Point{point.X * factor, point.Y * factor}
}

func ScalePointPointer(point *Point, factor int) {
	point.X *= factor
	point.Y *= factor
}

func main() {
	// Using order-based form
	// Here, we need to remember the order of the struct fields
	// and need to pass value for each and every struct field in order
	p := Point{1, 2}
	fmt.Println(p.X, p.Y) // 1 2

	// Using field keywords
	// Here, we need to pass the values of the fields that we want using
	// the name of the respected field. If the field is omitted, it will
	// be set to the zero value for its type
	// Because names are provided, the order of fields doesn't matter
	q := Point{X: 3}
	fmt.Println(q.X, q.Y) // 3 0

	// Passing struct to function
	scaledPoint := ScalePoint(p, 2)
	fmt.Println(scaledPoint.X, scaledPoint.Y) // 2 4

	// For efficiency, larger struct types are usually passed to or
	// returned from functions indirectly using a pointer
	newPoint := Point{4, 5}
	fmt.Println(newPoint.X, newPoint.Y) // 4 5
	ScalePointPointer(&newPoint, 2)
	fmt.Println(newPoint.X, newPoint.Y) // 8 10
}
