package main

import "fmt"

func push(stack []int, num int) []int {
	return append(stack, num)
}

func pop(stack []int) ([]int, int) {
	top := stack[len(stack)-1] // top of stack
	stack = stack[:len(stack)-1]
	return stack, top
}

func main() {
	var stack []int
	// Push elements to stack
	stack = push(stack, 1)
	stack = push(stack, 2)
	stack = push(stack, 3)
	// Pop elements from stack
	stack, popped_number := pop(stack)
	fmt.Printf("%d\n", popped_number) // "3"
	fmt.Printf("%d\n", stack)         // "[1 2]"
}
