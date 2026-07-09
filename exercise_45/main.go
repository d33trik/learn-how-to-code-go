package main

import "fmt"

func main() {
	a := func() {
		fmt.Println("I'm a func expression")
	}

	a()
}
