package main

import (
	"fmt"
	"sync"
)

// Run with the `-race` flag to prove that the race condition is fixed
func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	var counter int
	for range 100 {
		wg.Go(func() {
			m.Lock()
			v := counter
			v++
			counter = v
			m.Unlock()
		})
	}

	wg.Wait()
	fmt.Println(counter)
}
