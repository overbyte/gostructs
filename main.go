package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Cricket",
		contactInfo: contactInfo{
			email:   "jimminycrickey@gmail.com",
			zipCode: 12345,
		},
	}

	// convert jim into a memory address and save
	jimPointer := &jim
	// run updateName on the memory address of jim
	jimPointer.updateName("Jimminy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// use the memory address of type person as receiver
func (pointerToPerson *person) updateName(newFirstName string) {
	// dereference the memory address back to the value to update its members
	(*pointerToPerson).firstName = newFirstName
}
