package functional

// Person object
type Person struct {
	name string
	age  int
}

// NewPersonFactory will create new person in 2 steps. It will return
// function that takes age and then return Person struct
func NewPersonFactory(age int) func(name string) *Person {
	return func(name string) *Person {
		return &Person{name, age}
	}
}
