package main

import "fmt"

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Println("My name is", p.name)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{"James"}

	// You CANNOT pass a value of type person into saySomething
	// Cannot use p (variable of struct type Person) as Human value in argument to saySomething: Person does not implement Human (method speak has pointer receiver)
	// saySomething(p)

	// You CAN pass a value of type *person into saySomething
	saySomething(&p)

	// p is a value, but Go automatically takes its address
	// when calling a pointer receiver method on an addressable value
	p.speak()
}
