package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine engine
	make   string
	model  string
	color  string
	doors  int8
}

func main() {
	v1 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "Jeep",
		model: "Renegade",
		color: "Black",
		doors: 4,
	}

	fmt.Println(v1)

	v2 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "GWM",
		model: "Ora 5",
		color: "Silver",
		doors: 4,
	}

	fmt.Println(v2)
}
