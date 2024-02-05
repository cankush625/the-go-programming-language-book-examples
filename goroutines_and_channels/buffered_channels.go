package main

import "fmt"

func main() {
	// Create buffered channel of capacity 3
	ch := make(chan int, 3)
	// A send operation on a buffered channel inserts an element at the back of the
	// queue, and a receive operation removes an element from the front.
	// We can send upto 3 integers to the above channel without blocking the goroutine.
	// If the channel is full, the send operation blocks its goroutine until space
	// is made available by another goroutine's receive.
	ch <- 1
	ch <- 3
	ch <- 5
	// If channel is empty, a receive operation blocks until a value is sent by
	// another goroutine.

	// If program needs to know channel's buffer capacity, it can be obtained
	// by calling the built-in cap function
	fmt.Println(cap(ch)) // 3
	// To know the number of elements currently buffered, we can use the built-in
	// len function
	<-ch
	fmt.Println(len(ch)) // 2
}
