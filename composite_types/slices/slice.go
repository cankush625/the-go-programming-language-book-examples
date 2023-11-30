package main

import "fmt"

func main() {
	months := []string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}
	fmt.Println("Quarter 2: ", months[3:6]) // "Quarter 2:  [April May June]"

	// Slicing beyond capacity can cause a panic but
	// slicing beyond length extends the slice, so the
	// result may be longer than the original
	summer := months[2:6]
	fmt.Println(summer) // "[March April May June]"
	// Since 20 is out of capacity, this will cause panic
	// fmt.Println(summer[:20]) // panic: runtime error: slice bounds out of range [:20] with capacity 10
	extendedSummer := summer[:6]
	fmt.Println(extendedSummer) // "[March April May June July August]"

	// Note: Slice contains a pointer to the underlying array elements
	// So, passing a slice to the function and modifying that slice in
	// function will indeed modify the underlying array elements. And
	// hence original slice will also be modified.

	// If we need to test whenther a slice is empty, then user len(s),
	// and not s == nil. Because, there might a non-nil slice of length
	// and capacity zero, such as []int{} or make([]int, 3)

	// Note on slice comparison:
	// Comparing two slices is not simple, efficient and obvious.
	// So, the safest choice is to disallow slice comparisons altogether
	// The only legal slice comparison is against nil
	if summer == nil {
		/* ... */
	}
}
