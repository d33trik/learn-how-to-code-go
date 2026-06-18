package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for range 100 {
		if x := rand.Intn(5); x == 3 {
			fmt.Println("x is 3")
		}
	}
}
