package main

import (
	"log"
	"time"
)

func main() {
	// call newDeck method from deck module
	// cards := newDeck()
	cards, err := loadDeckFromFile("cardDeck.csv")
	if err != nil {
		log.Fatal(err)
	}
	//create random seed on each shuffle
	//using unix time in seconds sinds 1970
	seed := time.Now().Unix()
	cards.shuffle(seed)
	cards.print()
	// call print method defined in deck.go
	// cards.print()

	// deal the cards
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("cardDeck.csv")
}

// func newCard() string {
// 	return "Ace of Spades"
// }
