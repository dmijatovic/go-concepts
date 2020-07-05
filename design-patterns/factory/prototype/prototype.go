package prototype

import "errors"

const (
	//Adult constant
	Adult = iota
	//Child constant
	Child = iota
)

//Person is the struct
type Person struct {
	Name  string
	Adult bool
}

//NewPersonFactory takes role and returns function that accepts
//name and in its turn returns Person object/struct
func NewPersonFactory(role int) func(name string) *Person {
	return func(name string) *Person {
		switch role {
		case Adult:
			return &Person{Name: name, Adult: true}
		case Child:
			return &Person{Name: name, Adult: false}
		default:
			panic(errors.New("Unknon role type"))
		}
	}
}
