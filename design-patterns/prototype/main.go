package main

/* Prototype demo shows how to copy struct that uses pointers.
Method DeepCopy shows how to convert object
*/

import (
	"bytes"
	"encoding/gob"
	"log"
)

// Person is person with name and address
type Person struct {
	Name    string
	Address *Address
	Friends []string
}

// Address is persons address
type Address struct {
	Street, City, Country string
}

// DeepCopy will copy all data from person by marshaling object and creating a newone. We use binary serialization for this. In some other cases json serialization could also be an option.
func (p *Person) DeepCopy() *Person {
	//create buffer
	b := bytes.Buffer{}
	//create new encoder
	e := gob.NewEncoder(&b)
	//encode person into buffer b
	_ = e.Encode(p)

	// now decode into new object
	dc := gob.NewDecoder(&b)
	copy := Person{}
	_ = dc.Decode(&copy)
	// return copy
	return &copy
}

func wrongCopyOfStruct() {
	log.Println("Wrong copying object when struct users pointer definitions")

	dusan := Person{
		Name: "Dusan",
		Address: &Address{
			Street:  "Clauskinderweg",
			City:    "Amsterdam",
			Country: "Nederlands",
		},
	}
	// this does not work when we use pointers
	jane := dusan
	jane.Name = "Jane"
	jane.Address.City = "London"
	log.Println(dusan.Name, dusan.Address.City)
	log.Println(jane.Name, jane.Address.City)

}

func safeCopyOfObjects() {
	log.Println("Safe copying object")

	dusan := Person{
		Name: "Dusan",
		Address: &Address{
			Street:  "Clauskinderweg",
			City:    "Amsterdam",
			Country: "Nederlands",
		},
	}
	// custom helper methods that will clone project
	jane := dusan.DeepCopy()
	jane.Name = "Jane"
	jane.Address.City = "London"
	log.Println(dusan.Name, dusan.Address.City)
	log.Println(jane.Name, jane.Address.City)

}

func main() {
	wrongCopyOfStruct()
	safeCopyOfObjects()
}
