package main

import "fmt"

func main() {
	// Zero value
	var a int
	fmt.Println(a)

	// Shor declaration operator
	b := 42
	fmt.Println(b)

	// Multiple declaration
	c, d := 42, "no panic!"
	fmt.Println("Multiple declaration:", c, d)

	// Var for specificty
	var e float32 = 42.42
	fmt.Println(e)

	// Blank identifier
	_, f := 41, 42
	fmt.Println(f)
}
