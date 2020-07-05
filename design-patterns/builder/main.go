package main

/*
Functional person builder
This approach uses personBuilder to create person structure. Person builder takes array of actions. Each action has anonymous function that performs actuall work (personMethod).
The adventage of this approach is that additional actions can be added later and easily extend the person structure. In other words this approach is extensible and flexible.
Another feature implementated in this demo is fluid interface for me also know as `chaning`. The feature enables chaining the methods.
*/

import "log"

type person struct {
	name, position string
}

type personBuilder struct {
	actions []personMethod
}

type personMethod func(*person)

// with method that returns personBuilder for enabling `chanining`
func (pb *personBuilder) Called(name string) *personBuilder {
	// append action with modifying method
	pb.actions = append(pb.actions, func(p *person) {
		p.name = name
	})
	return pb
}

// with method that returns personBuilder for enabling `chanining`
func (pb *personBuilder) WorksAs(position string) *personBuilder {
	// append action with modifying method
	pb.actions = append(pb.actions, func(p *person) {
		p.position = position
	})
	return pb
}

//Create person and return pointer to it
func (pb *personBuilder) Create() *person {
	p := person{}
	for _, action := range pb.actions {
		action(&p)
	}
	return &p
}

func main() {
	log.Println("Start builder")
	pb := personBuilder{}
	person := pb.Called("Dusan").WorksAs("Developer").Create()
	log.Println(*person)
}
