package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Printf("x: %d\n", x)

	if x >= 0 && x <= 100 {
		fmt.Println("Between 0 and 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("Between 101 and 200")
	} else if x >= 201 && x <= 250 {
		fmt.Println("Between 201 and 250")
	} else {
		fmt.Println("More than 250")
	}
}
