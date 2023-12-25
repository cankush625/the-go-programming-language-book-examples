package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// Distance
// We can write a method for a type by using extra
// parameter which is called the method's receiver.
// In the method below, p is the receiver.
// The convention to name a receiver is by using the
// first letter of its type. For Point type its p.
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

// Distance
// Methods can be associated with any Go type and not
// just struct.
// Here, Path is a named slice type, but still we can
// define methods for it.
func (p Path) Distance() float64 {
	sum := 0.0
	for i := range p {
		if i > 0 {
			sum += p[i-1].Distance(p[i])
		}
	}
	return sum
}

func main() {
	p := Point{1, 2}
	q := Point{5, 7}
	fmt.Println(p.Distance(q)) // "6.40"

	r := Path{
		{1, 1},
		{7, 1},
		{7, 5},
		{1, 1},
	}
	fmt.Println(r.Distance()) // "17.21"
}
