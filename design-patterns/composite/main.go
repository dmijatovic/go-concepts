package main

import (
	"fmt"
	"log"
	"strings"
)

type Shape struct {
	Name, Color string
	Children    []Shape
}

func (s *Shape) Traverse(sb *strings.Builder, depth int) string {
	sb.WriteString(strings.Repeat("*", depth))

	if len(s.Color) > 0 {
		sb.WriteString(s.Color + " ")
		// sb.WriteRune(" ")
	}
	sb.WriteString(s.Name + "\n")
	// sb.WriteRune("\n")

	for _, child := range s.Children {
		child.Traverse(sb, depth+1)
	}

	return fmt.Sprint(sb)
}

func newSquare(color string) *Shape {
	return &Shape{Name: "Square", Color: color, Children: nil}
}

func newCircle(color string) *Shape {
	return &Shape{Name: "Circle", Color: color, Children: nil}
}

func main() {
	log.Println("Composite demo")
	dw := Shape{"My draving", "", nil}

	dw.Children = append(dw.Children, *newCircle("red"))
	dw.Children = append(dw.Children, *newSquare("green"))

	group := Shape{"Group1", "", nil}
	group.Children = append(group.Children, *newCircle("blue"))
	group.Children = append(group.Children, *newSquare("grey"))

	dw.Children = append(dw.Children, group)

	sb := dw.Traverse(&strings.Builder{}, 0)
	log.Println(sb)
}
