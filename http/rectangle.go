// Q: Take a look at the following code. Does the `rectangle` type satisfy the `shape` interface?

// A: No, because `rectangle`'s version of the `area` function returns a float64, but the `shape` interface expects a return of type `int`

package main

import "fmt"

type shape interface {
	area() int
}

type square struct {
	sideLength int
}

type rectangle struct {
	height float64
	width  float64
}

func (s square) area() int {
	return s.sideLength * s.sideLength
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func printArea(s shape) {
	fmt.Println(s.area())
}
