//Codul de mai jos va functiona deoarece
//interfata `shape` defineste o singura metoda numita `area()` ce returneaza un `int`
// Asta inseamna ca orice tip de date care are o metoda numita `area()` si returneaza un `int`
//satisface interfata `shape`

package main

import "fmt"

type shape interface {
	area() int
}

type square struct {
	sideLength int
}

func (s square) area() int {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.area())
}
