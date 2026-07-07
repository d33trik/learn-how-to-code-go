package main

import "fmt"

type person struct {
	name string
	age  uint
}

func (p person) speak() {
	fmt.Printf("My name is %s, and I'm %d years old\n", p.name, p.age)
}

func main() {
	p := person{
		name: "James Bond",
		age:  37,
	}

	p.speak()
}
