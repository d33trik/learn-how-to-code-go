package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Printf("x: %d\n", x)

	switch {
	case x >= 0 && x <= 100:
		fmt.Println("Between 0 and 100")
	case x >= 101 && x <= 200:
		fmt.Println("Between 101 and 200")
	case x >= 201 && x <= 250:
		fmt.Println("Between 201 and 250")
	default:
		fmt.Println("More than 250")
	}
}
