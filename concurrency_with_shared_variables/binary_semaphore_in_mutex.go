package main

// We can use a channel of capacity 1 to ensure that
// at most one goroutine accesses a shared variable at a time.
// A semaphore that counts only to 1 is called a binary semaphore.
// This pattern of mutual exclusion is so useful that it is supported
// directly by the Mutex type from the sync package. It's Lock method
// acquires the token and its Unlock method releases it.
var sema = make(chan struct{}, 1) // a binary semaphore guarding balance
var balance int

func Deposit(amount int) {
	sema <- struct{}{} // acquire token
	balance = balance + amount
	<-sema // release token
}

func Balance() int {
	sema <- struct{}{} // acquire token
	b := balance
	<-sema // release token
	return b
}

func main() {

}
