package main

import "fmt"

func main() {
	defer fmt.Println("This will run third")

	fmt.Println("This will run first")

	defer fmt.Println("This will run second")
}
