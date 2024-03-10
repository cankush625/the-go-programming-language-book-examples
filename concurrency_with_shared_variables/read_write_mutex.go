package main

import "sync"

// It is a special kind of lock that allows read-only
// operations to proceed in parallel with each other,
// but write operations to have fully exclusive access.
// This lock is called a multiple readers, single writer lock,
// and in Go it's provided by sync.RWMutex.
// An RWMutex requires more complex internal bookkeeping, so,
// it makes it slower than a regular mutex for uncontended locks.
var rwmu = sync.RWMutex{}
var myBalance int

func GetMyBalance() int {
	rwmu.RLock()
	defer rwmu.RUnlock()
	return myBalance
}

func main() {

}
