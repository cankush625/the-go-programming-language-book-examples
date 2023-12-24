package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(f)
}

// Defer for on enter and on exit functionality using
// anonymous function
func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the extra parenthesis
	// lots of work here
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

// Modify the return values of enclosing function using the
// deferred anonymous function
func double(num int) (result int) {
	return num * 2
}

func triple(num int) (res int) {
	defer func() { res += num }()
	return double(num)
}

func main() {
	// Note:
	// Defer defers the function call until the function
	// containing defer statement has finished
	// Any number of calls may be deferred; they are
	// executed in the reverse of the order in which they
	// were deferred.
	// Defer statement is often used with paired operations
	// like open and close, connect and disconnect, or
	// lock and unlock to ensure that resources are released
	// in all cases, no matter how complex the context flow.
	// The right place for a defer statement that releases a
	// resource is immediately after the resource has been
	// successfully acquired.

	// The defer statement cal also be used to pair `on entry`
	// and `on exit` actions when debugging a complex function.
	// The final parenthesis in the defer statement are important
	// here.
	bigSlowOperation()

	// Note:
	// The deferred anonymous function can even change the values
	// that the enclosing function returns to its caller
	fmt.Println(triple(5)) // "15"

	// Note:
	// Deferring in loops requires extra attention because they can
	// have unintentional side effects.
	// If we are handling a file and if we deferred the file.Close
	// call in for loop and the list of files is very large then
	// we might run out of file descriptors because the deferred
	// file.Close call will only be executed once the function is
	// finished.
	// The solution for this is to call file.Close in separate
	// function and call that function in the for loop. This way
	// the deferred call will be executed once that function
	// finished execution and hence every file will be closed as
	// soon as it's processed.
}

// Output:
// 2023/12/24 17:26:34 enter bigSlowOperation
// 2023/12/24 17:26:44 exit bigSlowOperation (10.001489709s)
