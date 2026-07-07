package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(foo(a...))
	fmt.Println(bar(a))
}

func foo(n ...int) int {
	t := 0
	for _, i := range n {
		t += i
	}

	return t
}

func bar(n []int) int {
	t := 0
	for _, i := range n {
		t += i
	}

	return t
}
