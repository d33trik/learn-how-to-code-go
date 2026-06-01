package main

import "fmt"

func main() {
	var a string = "don't panic!"
	fmt.Printf("Value: %s; Type: %T\n", a, a)

	var b int = 42
	fmt.Printf("Value: %d; Type: %T\n", b, b)

	var c float64 = 42.42
	fmt.Printf("Value: %f; Type: %T\n", c, c)
}
