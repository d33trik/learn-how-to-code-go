package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

func (s square) area() float64 {
	return s.length * s.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, float64(2))
}

type shape interface {
	area() float64
}

func printArea(s shape) {
	fmt.Printf("Area: %.2f\n", s.area())
}

func main() {
	square := square{
		length: 5,
		width:  8,
	}
	printArea(square)

	circle := circle{
		radius: 4,
	}
	printArea(circle)
}
