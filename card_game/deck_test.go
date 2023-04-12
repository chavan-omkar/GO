package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace Of Spades" {
		t.Errorf("Expecting Ace of Spade but got %v", d[0])
	}

	if d[len(d)-1] != "Four Of Club" {
		t.Errorf("Expecting Four of Club but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndReadFromFile(t *testing.T) {

	os.Remove("_decktesting")

	deck := newDeck()

	deck.saveToFile("_decktesting")

	loadDeck := readFromFile("_decktesting")

	//this is added to test crashing part
	if len(loadDeck) != 160 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadDeck))
	}

	os.Remove("_decktesting")
}
