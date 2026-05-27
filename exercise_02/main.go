package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

const (
	d = iota
	e
	f
)

func main() {
	fmt.Printf("a = %d\n", a)
	fmt.Printf("b = %d\n", b)
	fmt.Printf("c = %d\n", c)
	fmt.Printf("d = %d\n", d)
	fmt.Printf("e = %d\n", e)
	fmt.Printf("f = %d\n", f)
}
