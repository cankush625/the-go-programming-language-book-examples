package main

import "fmt"

func main() {
	// Create complex numbers
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)

	// Complex number arithmetic
	fmt.Println(x * y) // (-5+10i)

	// Print real and imaginary part of complex number
	fmt.Println("Real: ", real(x*y))      // Real:  -5
	fmt.Println("Imaginary: ", imag(x*y)) // Imaginary:  10
}
