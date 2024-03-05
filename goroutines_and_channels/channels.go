package main

func main() {
	ch := make(chan int) // ch has type "chan int"

	// Channel is connection between goroutines.
	// A channel is a communication mechanism that lets one
	// goroutine send values to another goroutine.

	// Send value to channel
	ch <- 10  // s send statement
	x := <-ch // a receive expression in an assignment statement
	<-ch      // a receive statement; result is discarded
	println(x)

	// We can use built-in close function to close the channel
	close(ch)

	// A channel created with simple call to make is called an
	// unbuffered channel, but make accepts an optional second
	// argument, an integer called the channel's capacity.
	// If the capacity is non-zero, make creates a buffered channel.
	ch = make(chan int)    // unbuffered channel
	ch = make(chan int, 0) // unbuffered channel
	ch = make(chan int, 3) // buffered channel with capacity 3
}
