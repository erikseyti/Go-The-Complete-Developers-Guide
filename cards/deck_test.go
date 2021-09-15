package main

import (
	"os"
	"strconv"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 54 {
		// t.Errorf("test")

		t.Error("Expected deck lenght of 54, but got " + strconv.Itoa(len(d)))
		t.Errorf("Expected deck lenght of 54, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Red Joker" {
		t.Errorf("Expected last card to be Red Joker, but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	cleanDeckFromFile("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 54 {
		t.Errorf("Expected 54 cards in deck, got %v", len(loadedDeck))
	}
	cleanDeckFromFile("_decktesting")
}

func cleanDeckFromFile(filename string) {
	os.Remove(filename)
}
