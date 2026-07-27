package main

import (
	"errors"
	"log"
)

var errFoo = errors.New("custo error")

func main() {
	err := foo()
	if err != nil {
		log.Fatalln(err)
	}
}

func foo() error {
	return errFoo
}
