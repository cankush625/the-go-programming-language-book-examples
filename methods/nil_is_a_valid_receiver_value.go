package main

import "fmt"

// An IntList is linked list of integers
// A nil *IntList represents the empty list
type IntList struct {
	Value int
	Tail  *IntList
}

// Sum returns the sum of the list elements
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

func main() {
	list2 := IntList{2, nil}
	list1 := IntList{5, &list2}
	list := IntList{10, &list1}
	fmt.Println(list.Sum()) // "17"
}
