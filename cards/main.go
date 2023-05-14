package main

import "fmt"

func main() {

	// var card string = "Ace of spades"
	// alternative syntax for the above. Go will infer the type string
	// the colon before equals is only used for new variable initializations not for assignment
	card := "Ace of Spades"
	card = "Five of Diamonds"
	fmt.Println((card))
}
