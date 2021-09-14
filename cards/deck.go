package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	err := ioutil.WriteFile(filename, []byte(d.toString()), 0666)
	return err
}

func newDeckFromFile(filename string) deck {
	bs, err :=ioutil.ReadFile(filename)

	if err != nil {
		// How to handle a error?
		// Option #1 - log the file and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error: ",err)
		os.Exit(1)
	}

	s :=strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r :=rand.New(source)


	for i := range d {
		newPosition := r.Intn(len(d)-1)

		d[i], d[newPosition] = d[newPosition] , d[i]
	}
}