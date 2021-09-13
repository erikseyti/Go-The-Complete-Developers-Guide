package main

import "fmt"

// Create a new type of 'deck'
// whic is a slice of strings
type deck []string

// This is called Receiver Functions
// with the ( d deck ) before the declaration of the print() function, any instance of the type deck have the acess to this function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three",
		"Four", "Five", "Six", "Seven", "Eight",
		"Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	cards = append(cards, "Black Joker")
	cards = append(cards, "Red Joker")

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) addFiveOfDiamonds() deck {
	d = append(d, "Five of Diamonds")
	return d
}
