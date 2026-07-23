package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Run with the `-race` flag to prove the race condition
func main() {
	var wg sync.WaitGroup

	var counter int
	for range 100 {
		wg.Go(func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
		})
	}

	wg.Wait()
	fmt.Println(counter)
}
