package main

import "fmt"

func main() {

	// var card string = "Ace of Spades"
	// alternative syntax for the above. Go will infer the type string
	// the colon before equals is only used for new variable initializations not for assignment
	card := newCard()

	fmt.Println((card))
}

func newCard() string {
	return "Five of Diamonds"
}
