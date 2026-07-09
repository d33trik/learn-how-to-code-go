package main

import "fmt"

func main() {
	x := printSquare(square, 2)

	fmt.Println(x)
}

func square(n int) int {
	return n * n
}

func printSquare(callback func(int) int, n int) string {
	return fmt.Sprintf("Square: %d", callback(n))
}
