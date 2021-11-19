package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}
type portugueseBot struct {}

type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	ptb := portugueseBot{}

	printGreeting(eb)
	printGreeting(sb)
	printGreeting(ptb)
	botTypeMessage(ptb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func botTypeMessage(b bot) {
	fmt.Println("I am using a function receiver from type bot!")
}

func (englishBot) getGreeting() string {
	// Very custom logic for generating an english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func (portugueseBot) getGreeting() string {
	return "Ola!"
}