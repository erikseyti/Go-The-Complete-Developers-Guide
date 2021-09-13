package main

import "fmt"

func main() {
	cards := deck{newCard(), "Ace of Spades"}

	cards = append(cards, "Six of Spades")

	fmt.Println(cards)

	// cards = cards.addFiveOfDiamonds()

	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}

func newNumberCard() int {
	return 1
}
