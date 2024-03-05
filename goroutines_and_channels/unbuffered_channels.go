package main

import (
	"io"
	"log"
	"net"
	"os"
)

// A send operation on unbuffered channel blocks the sending goroutine until another
// goroutine executes a corresponding receive on the same channel, at which point
// the value is transmitted and both goroutines may continue.
// Conversely, if the receive operation was attempted first, the receiving goroutine
// is blocked until another goroutine performs a send on the same channel.

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done
}
