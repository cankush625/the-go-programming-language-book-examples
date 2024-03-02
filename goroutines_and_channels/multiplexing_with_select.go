package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // Read a single byte
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown. Press return to abort.")
	select {
	case count := <-time.After(10 * time.Second):
		// Do nothing
		fmt.Println("On track ", count) // "On track  2024-03-02 22:49:34.758271 +0530 IST m=+10.001181709"
	case <-abort:
		fmt.Println("Mission aborted!")
		return
	}
	fmt.Println("Mission started!")
}
