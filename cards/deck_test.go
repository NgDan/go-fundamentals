package main

import (
	"os"
	"testing"
)

// in go, test files simply have the _test suffix

// convention in go is to PascalCase the test functions
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of16, but got %v", len(d))
		// we could also use Error instead of Errorf
		// t.Error("Expected deck length of 16, but got", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAdNewDeckTestFile(t *testing.T) {
	// very important to do the cleanup at the beginning
	// of the test to always make sure we start on a
	// clean slate
	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	deckFromFile := newDeckFromFile("_decktesting")

	if len(deckFromFile) != 16 {
		t.Errorf("Expected deck length to be 16 but got %v", len(deckFromFile))
	}

	os.Remove("_decktesting")
}
