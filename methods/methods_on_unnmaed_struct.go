package main

import (
	"fmt"
	"sync"
)

// We can do embeddings on unnamed struct type
// So, unnamed struct type can have methods
var cache = struct {
	sync.Mutex
	mapping map[string]int
}{
	mapping: make(map[string]int),
}

// Using embedding on unnamed struct, we have more
// expressive variable names to the variables related
// to the cache
func Lookup(key string) int {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

func main() {
	cache.mapping["views"] = 10
	fmt.Println(Lookup("views")) // "10"
}
