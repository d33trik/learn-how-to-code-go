package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Run with the `-race` flag to prove that the race condition is fixed
func main() {
	var wg sync.WaitGroup

	var counter int64
	for range 100 {
		wg.Go(func() {
			atomic.AddInt64(&counter, 1)
		})
	}

	wg.Wait()
	fmt.Println(counter)
}
