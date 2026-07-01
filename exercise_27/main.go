package main

import "fmt"

func main() {
	states := []string{
		`Alabama`,
		`Alaska`,
		`Arizona`,
		`Arkansas`,
		`California`,
		`Colorado`,
		`Connecticut`,
		`Delaware`,
		`Florida`,
		`Georgia`,
		`Hawaii`,
		`Idaho`,
		`Illinois`,
		`Indiana`,
		`Iowa`,
		`Kansas`,
		`Kentucky`,
		`Louisiana`,
		`Maine`,
		`Maryland`,
		`Massachusetts`,
		`Michigan`,
		`Minnesota`,
		`Mississippi`,
		`Missouri`,
		`Montana`,
		`Nebraska`,
		`Nevada`,
		`New Hampshire`,
		`New Jersey`,
		`New Mexico`,
		`New York`,
		`North Carolina`,
		`North Dakota`,
		`Ohio`,
		`Oklahoma`,
		`Oregon`,
		`Pennsylvania`,
		`Rhode Island`,
		`South Carolina`,
		`South Dakota`,
		`Tennessee`,
		`Texas`,
		`Utah`,
		`Vermont`,
		`Virginia`,
		`Washington`,
		`West Virginia`,
		`Wisconsin`,
		`Wyoming`,
	}

	x := make([]string, 0, len(states))
	x = append(x, states...)

	fmt.Printf("len(x): %v\n", len(x))
	fmt.Printf("cap(x): %v\n", cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Printf("[%d]: %s\n", i, x[i])
	}
}
