package main

import "fmt"

type shape interface {
	area() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	sq := square{2.0}
	tr := triangle{3.0, 2.0}

	printArea(sq)
	printArea(tr)

}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.area())
}
