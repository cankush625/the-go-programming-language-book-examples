package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Create array of length 3 and initialized to zeroth
	// value of it's type
	var a [3]int
	fmt.Println(a) // "[0, 0, 0]"
	// Print first element of array
	fmt.Println(a[0]) // "0"
	// Print last element of array
	fmt.Println(a[len(a)-1]) // "0"
	// Get type of array a
	// Array type always contains the length of array
	fmt.Println(reflect.TypeOf(a)) // "[3]int"
	// Create array with 1, 2, 3 numbers
	var b = [3]int{1, 2, 3}
	fmt.Println(b) // "[1, 2, 3]"
	// If any value is missing in literal at a position then
	// that position will have a zeroth value of the array type
	var c = [3]int{1, 2}
	fmt.Println(c) // "[1, 2, 0]"
	// If ellipsis '...' appears in place of length, then
	// array length is determined by number of initializers
	var d = [...]int{1, 2, 3, 4}
	fmt.Println(len(d)) // "4"
}
