package main

import "fmt"

func main() {
	a := foo()
	b, c := bar()

	fmt.Println(a)
	fmt.Println(b, c)
}

func foo() int {
	return 42
}

func bar() (string, int) {
	return "Answer to the Ultimate Question of Life, the Universe, and Everything", 42
}
