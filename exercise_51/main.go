package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Printf("a value: %v\n", a)
	fmt.Printf("a type: %T\n", a)
	fmt.Printf("a data: %v\n", *a)

	fmt.Println()

	fmt.Printf("b value: %v\n", b)
	fmt.Printf("b type: %T\n", b)
	fmt.Printf("b data: %v\n", *b)

	fmt.Println()

	fmt.Printf("c value: %v\n", c)
	fmt.Printf("c type: %T\n", c)
	fmt.Printf("c data: %v\n", *c)

	fmt.Println()

	fmt.Printf("d value: %v\n", d)
	fmt.Printf("d type: %T\n", d)
	fmt.Printf("d data: %v\n", *d)
}
