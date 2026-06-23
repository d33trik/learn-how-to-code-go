package main

import "fmt"

func main() {
	a := [5]int{}

	for i := range cap(a) {
		a[i] = i
	}

	for i, v := range a {
		fmt.Printf("i: %d; v: %d\n", i, v)
	}
}
