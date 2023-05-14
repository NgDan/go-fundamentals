package main

import "fmt"

// create a new type of deck which is a slice of strings

type deck []string

// This syntax (d deck) means that any variable of type deck
// will now have access to this print method
// This is a receiver function.
// The "d" is the actual working variable of type deck
// that will be passed to the function
// By convention the name of the variable is the first
// or first two letters of the type. In this case "d"
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// this function doesn't need a receiver because it
// is not meant to be related to any type, it'll
// just return a new deck of type deck
func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, cardSuit := range cardSuites {
		for _, cardValue := range cardValues {

			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}
	return cards
}
