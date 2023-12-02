package main

import "fmt"

func appendInt(slice []int, nums ...int) []int {
	var newSlice []int
	newSliceLen := len(slice) + len(nums)
	if newSliceLen <= cap(slice) {
		// There is room to grow. Extend the slice.
		newSlice = slice[:newSliceLen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		newSliceCap := newSliceLen
		if newSliceCap < 2*len(slice) {
			newSliceCap = 2 * len(slice)
		}
		newSlice = make([]int, newSliceLen, newSliceCap)
		copy(newSlice, slice)
	}
	// expand newSlice to at least newSliceLen
	copy(newSlice[len(slice):], nums)
	return newSlice
}

func main() {
	// Built-in append function appends items to slices
	var runes []rune
	for _, r := range "Hello World" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o' ' ' 'W' 'o' 'r' 'l' 'd']"

	// Go built-in conversion technique
	var newRunes = []rune("Hello World")
	fmt.Printf("%q\n", newRunes) // "['H' 'e' 'l' 'l' 'o' ' ' 'W' 'o' 'r' 'l' 'd']"

	// Using the append function we have created
	// Refer page 89 in the book for detailed explanation
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = appendInt(numbers, i)
	}
	fmt.Printf("%d\n", numbers) // "[0 1 2 3 4 5 6 7 8 9]"

	// Append multiple numbers to slice at once
	digits := []int{1, 2, 3}
	digits = appendInt(digits, 4, 5)
	fmt.Printf("%d\n", digits) // "[1 2 3 4 5]"

	// Note:
	// We should always assign the output of append function
	// to the original slice variable. We must not assume that any
	// operations made on older slice will be reflected in new slice.
	// So, the result of the call to append function should be assigned
	// to the same slice variable.
}
