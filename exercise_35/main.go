package main

import "fmt"

func main() {
	p := struct {
		name           string
		friends        map[string]int
		favoriteDrinks []string
	}{
		name: "James Bond",
		friends: map[string]int{
			"Jenny": 27,
			"Q":     87,
			"Ian":   47,
		},
		favoriteDrinks: []string{
			"Martini",
			"Water",
		},
	}

	fmt.Println(p)
}
