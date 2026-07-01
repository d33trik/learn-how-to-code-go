package main

import "fmt"

func main() {
	w := make(map[string]int)

	for _, v := range words {
		w[v]++
	}

	for k, v := range w {
		fmt.Printf("%s: %d\n", k, v)
	}
}
