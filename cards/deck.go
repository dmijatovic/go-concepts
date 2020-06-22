package main

import "fmt"

// declare deck as array of strings
type deck []string

// create new deck with all cards
func newDeck() deck {
	//create empty new deck
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Carot"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, val := range cardValues {
			cards = append(cards, val+" of "+suit)
		}
	}

	return cards
}

/*
	Deal function will return 2 values
*/
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// declare receiver function
// any variable of type deck
// gets access to print method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
