package main

func main() {
	// cards := newDeck()

	// // Deal returns two values. The first one will be
	// // assigned to hand and the second one to remainingDeck
	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()

	cards := newDeck()
	cards.saveToFile("cards")

	exampleDeck := newDeckFromFile("cards")
	exampleDeck.print()
}
