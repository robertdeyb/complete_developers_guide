package main

import "fmt"

func main() {
	t := triangle{base: 2, height: 5}
	s := square{sideLength: 5}

	printArea(t)
	printArea(s)
}

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
