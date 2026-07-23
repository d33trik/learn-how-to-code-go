package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("foo")
		wg.Done()
	}()

	wg.Go(func() {
		fmt.Println("bar")
	})

	wg.Wait()
}
