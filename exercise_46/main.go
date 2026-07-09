package main

import "fmt"

func main() {
	r := deepThoughtResponse()

	fmt.Println(r())
}

func deepThoughtResponse() func() int {
	return func() int {
		return 42
	}
}
