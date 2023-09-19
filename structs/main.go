package main

//person             contactInfo
//|firstname  |      |email  |
//|lastname   |      |zipCode|
//|contactInfo|

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string //ordinea conteaza intotdeauna
	lastName  string
	contactInfo
}

func main() {
	//method 1
	// p1 := person{"Alex", "Anderson"} //ordinea trebuie sa fie *EXACT* cum a fost definit structul
	//method 2
	// p2 := person{firstName: "Alex", lastName: "Anderson"}
	//fmt.Println(p1, p2)
	//un alt print
	//fmt.Printf("%+v", p1) //print: {firstName:Alex lastName:Anderson}

	//syntax for updating a struct
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jimPointer := &jim //adresa din RAM
	jimPointer.updateName("Jimmy")
	//jim.print()
	//jim.updateName("Jimmy")
	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
