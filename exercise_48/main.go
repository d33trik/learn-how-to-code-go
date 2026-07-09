package main

import "fmt"

func main() {
	x := counter()

	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}

func counter() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}
