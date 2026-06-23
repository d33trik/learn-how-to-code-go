package main

import "fmt"

func main() {
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Printf("s[]: %v\n", s)
	fmt.Printf("s[:5]: %v\n", s[:5])
	fmt.Printf("s[5:]: %v\n", s[5:])
	fmt.Printf("s[2:7]: %v\n", s[2:7])
	fmt.Printf("s[1:6]: %v\n", s[1:6])
}
