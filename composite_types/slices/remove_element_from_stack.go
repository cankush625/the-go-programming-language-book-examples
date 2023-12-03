package main

import "fmt"

func remove_by_preserving_order(slice []int, position int) []int {
	// To preserve the order after removing the number
	// slide the higher-numbered elements down by one to fill
	// the gap
	copy(slice[position:], slice[position+1:])
	return slice[:len(slice)-1]
}

func remove_without_preserving_order(slice []int, position int) []int {
	// Move the last element to the place of element which
	// needs to be removed
	slice[position] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	s := []int{4, 5, 6, 7, 8}
	fmt.Println(remove_by_preserving_order(s, 2)) // "[4 5 7 8]"

	s = []int{4, 5, 6, 7, 8}
	fmt.Println(remove_without_preserving_order(s, 2)) // "[4 5 8 7]"
}
