package main

import "fmt"

// Variadic function is a function that accepts varying number of arguments
// To declare variadic function, the type of final parameter
// is preceded by an ellipsis, "...", which indicates that the function
// may be called with any number od arguments of this type.
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func f(...int) {}
func g([]int)  {}

func main() {
	fmt.Println(sum())           // "0"
	fmt.Println(sum(1))          // "1"
	fmt.Println(sum(1, 2, 3, 4)) // "10"
	values := []int{1, 2, 3, 4}
	// By placing ellipsis after the final argument we can pass
	// list of arguments if the final argument is a slice
	fmt.Println(sum(values...)) // "10"

	// Note:
	// Although ...int behaves like a slice within the function body,
	// the type of variadic function is distinct from the type
	// of a function with an ordinary slice parameter.
	fmt.Printf("%T\n", f) // "func(...int)"
	fmt.Printf("%T\n", g) // "func([]int)"
}
