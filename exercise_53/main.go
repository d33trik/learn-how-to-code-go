package main

import "fmt"

type person struct {
	name string
}

func changeNameByValue(p person, n string) person {
	p.name = n
	return p
}

func changeNameByPointer(p *person, n string) {
	p.name = n
}

func main() {
	p1 := person{
		name: "James",
	}

	fmt.Println(p1)
	p1 = changeNameByValue(p1, "Jenny")
	fmt.Println(p1)

	p2 := person{
		name: "Jenny",
	}

	fmt.Println(p2)
	changeNameByPointer(&p2, "James")
	fmt.Println(p2)
}
