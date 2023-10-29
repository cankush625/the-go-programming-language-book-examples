package main

import (
	"fmt"
	"math"
)

func main() {
	println("Max float value: ", math.MaxFloat32)

	// Smallest positive integer that cannot be exactly
	// represented as a float32 is not large:
	var f float32 = 16777216
	fmt.Println(f == f+1) // "true"
}
