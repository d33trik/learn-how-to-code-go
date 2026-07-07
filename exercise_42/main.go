package main

import "fmt"

func main() {
	a := doMath(42, 16, add)
	fmt.Println(a)

	b := doMath(42, 16, subtract)
	fmt.Println(b)
}

func doMath(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}
