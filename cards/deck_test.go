package main

import (
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