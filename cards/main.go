package main

func main() {
	// call newDeck method from deck module
	cards := newDeck()
	// call print method defined in deck.go
	// cards.print()

	// deal the cards
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}

func newCard() string {
	return "Ace of Spades"
}
