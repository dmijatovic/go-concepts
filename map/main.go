package main

import "fmt"

func main() {
	//basic declaration
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#4bf745",
		"blue":  "#ef45fgf",
	}

	//using var
	// var colors2 map[string]string
	// using make command
	colors3 := make(map[int]string)

	// add key
	colors3[10] = "Value 1"
	//delete key
	delete(colors3, 10)

	// fmt.Println(colors, colors2, colors3)
	printMap(colors)
}

func printMap(c map[string]string) {
	for key, val := range c {
		fmt.Println(key, val)
	}
}
