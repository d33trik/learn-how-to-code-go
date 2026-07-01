package main

import "fmt"

func main() {
	x := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "I'm 008"},
	}

	for _, i := range x {
		for _, j := range i {
			fmt.Println(j)
		}
	}
}
