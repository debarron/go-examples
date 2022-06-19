package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLegth float64
}
type triangle struct {
	height float64
	base   float64
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func (s square) getArea() float64 {
	return s.sideLegth * s.sideLegth
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	triangleEx := triangle{height: 3.12, base: 12.3}
	squareEx := square{sideLegth: 34.2}

	printArea(triangleEx)
	printArea(squareEx)
}
