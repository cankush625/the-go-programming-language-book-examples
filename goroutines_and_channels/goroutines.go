package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func main() {
	// Creates a separate goroutine that runs independently.
	// The go statement itself completes immediately and
	// causes the function followed by it to run in the
	// newly created goroutine.
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // Slow
	fmt.Printf("\rfibonacci(%d) = %d\n", n, fibN)
}
