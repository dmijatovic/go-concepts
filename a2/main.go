package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLenght float64
}

type triangle struct {
	height float64
	base   float64
}

func (s square) getArea() float64 {
	return s.sideLenght * s.sideLenght
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(sh shape) {
	val := sh.getArea()
	fmt.Println(val)
}

func main() {
	println("It works")

	sq := square{sideLenght: 10}
	tr := triangle{height: 10, base: 10}

	printArea(sq)
	printArea(tr)
}
