package main

import (
	"fmt"
	"math"
)

func main() {
	// math.Pi is an untyped constant, so it can be used
	// wherever any floating point or complex value is needed
	// Otherwise it would have required explicit type conversion
	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi

	fmt.Println(x, y, z) // "3.1415927 3.141592653589793 (3.141592653589793+0i)"
}
