package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
)

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

/*
	Convert desck to string ; separated
*/
func (d deck) toString() string {
	return strings.Join([]string(d), ";")
}

func (d deck) saveToFile(fileName string) error {
	// convert deck into byte slice
	bytes := []byte(d.toString())
	// write to file
	err := ioutil.WriteFile(fileName, bytes, 0644)
	// check for errors
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func loadDeckFromFile(fileName string) (deck, error) {
	// read bytes from file
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		return newDeck(), err
	}
	//converts bytes to string
	tekst := string(bs)
	// split to array
	array := strings.Split(tekst, ";")
	// create deck
	cards := deck(array)
	// return cards
	return cards, nil
}

func (d deck) shuffle(seed int64) {
	src := rand.NewSource(seed)
	r := rand.New(src)
	r.Shuffle(len(d), func(i, j int) {
		println(i, j)
		d[i], d[j] = d[j], d[i]
	})
}

// declare receiver function
// any variable of type deck
// gets access to print method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
