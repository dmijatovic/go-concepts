package main

import (
	"fmt"
)

type person struct {
	id        string
	firstName string
	lastName  string
	contact   contactInfo
	// birthDate struct time.Date
}

type contactInfo struct {
	email   string
	phone   string
	address string
}

func main() {
	// declare by position
	alex := person{"1234", "Alex", "Krimbeart", contactInfo{"email", "phone", "one road"}}
	// declaring by key
	dusan := person{id: "2345", firstName: "Dussan", lastName: "Mijatovic",
		contact: contactInfo{
			email:   "email@com",
			phone:   "myphone",
			address: "my address"}}
	// declaring empty struct (will have initial default values)
	var marco person
	// assigning values
	marco.firstName = "Marco"
	marco.lastName = "Markovic"

	fmt.Println("It works...", alex, dusan, marco)
	// create pointer to struct marco
	// marcoPointer := &marco
	// pass the pointer
	// marcoPointer.updateName("Marcius")
	marco.updateName("Markus")
	//print changed name on marco struct
	marco.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// call the function with pointer to person struct
// go wil pass the pointer instead of value automatically
func (pointer *person) updateName(fn string) {
	(*pointer).firstName = fn
}
