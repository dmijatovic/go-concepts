package factory

/* Interface factory which creates person but protects struct/object from being manipulated from outside. The function NewPerson returns interface instead of struct directly. This means that we can return different types of object and methods depending on the intialization.

In this example we return adult not adult person method SayHello.
*/

import "log"

// Person interface
type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

type adult struct {
	name string
	age  int
}

//NewPerson returns interface which in turn returns 'protected' struct
//of a person. Because we return interface we do not need * in declaration
//BUT we need to return pointer to internal person struct.
func NewPerson(name string, age int) Person {
	if age > 17 {
		return &adult{name, age}
	} else {
		return &person{name, age}
	}
}

func (p *person) SayHello() {
	log.Println("Hello ", p.name, ". Sory but you are not adult")
}

func (p *adult) SayHello() {
	log.Println("Hello ", p.name, ". You are adult!")
}
