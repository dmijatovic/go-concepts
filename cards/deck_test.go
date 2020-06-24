package main

import (
	"fmt"
	"os"
	"testing"
)

func checkDeckLength(d deck) (ok bool, err string) {
	length := 16
	if len(d) != length {
		message := fmt.Sprintf("Expected deck length of %d but got %v", length, len(d))
		// t.Errorf()
		return false, message
	}
	return true, ""
}

func TestNewDeck(t *testing.T) {
	d := newDeck()

	ok, err := checkDeckLength(d)

	if ok == false {
		t.Errorf(err)
	}
	firstCard := "Ace of Spades"
	if d[0] != firstCard {
		t.Errorf("Expected first card to be %v but got %v", firstCard, d[0])
	}
	lastCard := "Four of Carot"
	if d[len(d)-1] != lastCard {
		t.Errorf("Expected last cart to be %v but got %v", lastCard, d[len(d)-1])
	}
}

func TestSaveToAndLoadFromFile(t *testing.T) {
	fileName := "_decktesting"
	os.Remove(fileName)

	cards := newDeck()
	err := cards.saveToFile(fileName)

	cardsFromFile, err := loadDeckFromFile(fileName)

	if err != nil {
		t.Errorf("Failed to load from file: %v", err)
	}

	ok, msg := checkDeckLength(cardsFromFile)

	if ok == false {
		t.Errorf(msg)
	}

	os.Remove(fileName)
}
