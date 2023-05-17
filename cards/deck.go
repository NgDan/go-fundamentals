package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

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

// the (deck, deck) return type tells go we're going to return 2 values
// both of type deck
func deal(d deck, handSize int) (deck, deck) {

	// to return 2 return values we simply put a comma between them
	// this d[startIndexIncluding:upToNotIncluding] syntax is
	// how we create a subset of a slice based on start (includint)
	// and end index (NOT including)
	return d[:handSize], d[handSize:]
}

// we'll make this function get a receiver of type deck
// to tie it to the deck type
func (d deck) toString() string {
	// this is called type conversion. Here we convert the value
	// d of type deck to a slice of strings with the type slice
	// of strings
	sliceOfStrings := []string(d)
	joinedSliceOfStrings := strings.Join(sliceOfStrings, ",")
	return joinedSliceOfStrings
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) readFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		// exit current program with given status code
		os.Exit(1)
	}
	deckString := string(byteSlice)

	sliceOfStrings := strings.Split(deckString, ",")

	// type conversion from sliceOfStrings to a deck.
	// the difference between deck and []string is that
	// deck has these additional functions attached to
	// it declared by us using a receiver
	return deck(sliceOfStrings)
}

func (d deck) shuffle() {
	for i := range d {

		// get a pseudo-random number between 0 and length of the
		// d slice minus 1
		newPosition := rand.Intn(len(d) - 1)

		// this syntax swaps elements in the slice
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
