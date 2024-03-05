package main

import "fmt"

// Unidirectional channels are used for only one use case.
// They can either be sent to or received from.

// Counter
func counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out)
}

// Squarer
func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

// Output:
// 0
// 1
// 4
// 9
// 16
// 25
// 36
// 49
// 64
// 81
