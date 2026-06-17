package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47}

	for i, v := range x {
		fmt.Printf("index: %d; value: %d\n", i, v)
	}
}
