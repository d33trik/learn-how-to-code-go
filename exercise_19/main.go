package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	if james, ok := m["James"]; ok {
		fmt.Printf("James age is: %d\n", james)
	} else {
		fmt.Println("key not found")
	}

	if q, ok := m["Q"]; ok {
		fmt.Printf("Q age is: %d\n", q)
	} else {
		fmt.Println("key not found")
	}
}
