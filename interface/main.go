package main

import "fmt"

type bot interface {
	// all struct having this
	// func implemented
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	println("Test this")
	// example of basic
	// bot interface
	// eb := englishBot{}
	// sb := spanishBot{}
	// printGreeting(eb)
	// printGreeting(sb)

	// example interface from Go
	mainHTTP()
}
