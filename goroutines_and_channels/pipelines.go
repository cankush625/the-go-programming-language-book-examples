package main

import "fmt"

// Channels can be used to connect goroutines together so that the
// output of one is the input to another. This is called a pipeline.
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break // Channel was closed and drained
			}
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for {
		s, ok := <-squares
		if !ok {
			break
		}
		fmt.Println(s)
	}
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
