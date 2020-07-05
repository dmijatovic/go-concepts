package main

/*
 Factory function. Useful for default values and validation of initial values.
*/
import (
	"errors"
	"log"

	"./functional"
	factory "./interface"
	"./prototype"
)

// Adult person  age > 17
type Adult struct {
	Name         string
	Age          int
	NumberOfLegs int
}

// NewPerson creates adult person and returns pointer to values
func NewPerson(name string, age int) (*Adult, error) {
	if age > 17 {
		return &Adult{
			Name:         name,
			Age:          age,
			NumberOfLegs: 2,
		}, nil
	} else {
		return &Adult{}, errors.New("Person is not adult")
	}
}

func simpleFactory() {
	log.Println("Factory function")
	p1, e1 := NewPerson("Dusan", 49)
	if e1 != nil {
		log.Println(e1)
	}
	log.Println(*p1)
	p2, e2 := NewPerson("Jasa", 16)
	if e2 != nil {
		log.Println(e2)
	}
	log.Println(*p2)
}

func interfaceFactory() {
	p1 := factory.NewPerson("Dusan", 35)
	p1.SayHello()
	p2 := factory.NewPerson("Maria", 15)
	p2.SayHello()
}

func higherOrgerFunction() {
	adultsFactory := functional.NewPersonFactory(20)
	dusan := adultsFactory("Dusan")
	jongsFactory := functional.NewPersonFactory(16)
	log.Println(*dusan)
	maria := jongsFactory("Maria")
	log.Println(*maria)
}

func prototypeFactory() {
	adultsFactory := prototype.NewPersonFactory(prototype.Adult)
	dusan := adultsFactory("Dusan")
	childFactory := functional.NewPersonFactory(prototype.Child)
	log.Println(*dusan)
	maria := childFactory("Maria")
	log.Println(*maria)
}

func main() {
	simpleFactory()
	interfaceFactory()
	higherOrgerFunction()
	prototypeFactory()
}
