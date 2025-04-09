package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

func calcArea(s shape) float64 {
	return s.area()
}

type rectangle struct {
	width float64
	height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func main() {
	r := rectangle{10, 5}
	fmt.Println("This is calc of rectangle:", calcArea(r))

	c := circle{5}
	fmt.Println("This is calc of circle:", calcArea(c))
}