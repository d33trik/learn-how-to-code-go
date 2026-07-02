package main

import "fmt"

type person struct {
	firstName               string
	lastName                string
	favoriteIceCreamFlavors []string
}

func main() {
	p1 := person{
		firstName:               "Michael",
		lastName:                "Scott",
		favoriteIceCreamFlavors: []string{"Ice Mint", "Chocolate"},
	}

	fmt.Printf("First Name: %s\n", p1.firstName)
	fmt.Printf("Last Name: %s\n", p1.lastName)
	fmt.Println("Favorite Ice Cream Flavors:")
	for _, v := range p1.favoriteIceCreamFlavors {
		fmt.Printf("  - %s\n", v)
	}

	fmt.Println()

	p2 := person{
		firstName:               "Dwight",
		lastName:                "Schrute",
		favoriteIceCreamFlavors: []string{"Beets", "Mayonnaise and Black Olives"},
	}

	fmt.Printf("First Name: %s\n", p2.firstName)
	fmt.Printf("Last Name: %s\n", p2.lastName)
	fmt.Println("Favorite Ice Cream Flavors:")
	for _, v := range p2.favoriteIceCreamFlavors {
		fmt.Printf("  - %s\n", v)
	}
}
