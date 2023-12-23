package deck

import (
	"os"
	"slices"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", deck[0])
	}

	if deck[len(deck)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", deck[len(deck)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	deck := NewDeck()
	deck.SaveToFile()

	if _, err := os.Stat(store_filename); os.IsNotExist(err) {
		t.Errorf("Expected file %v to exist, but it did not", store_filename)
	}

	loadedDeck := NewDeckFromFile()

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(loadedDeck))
	}

	if slices.Compare(deck, loadedDeck) != 0 {
		t.Errorf("Expected decks to be equal, but they were not")
	}

	t.Cleanup(func() {
		os.Remove(store_filename)
	})
}

func TestDeal(t *testing.T) {
	deck := NewDeck()

	dealt, _ := deck.Deal(5)

	if len(dealt) != 5 {
		t.Errorf("Expected deck length of 5, but got %v", len(dealt))
	}

	if len(deck) != 47 {
		t.Errorf("Expected deck length of 47, but got %v", len(deck))
	}

	_, err := deck.Deal(50)

	if err == nil {
		t.Errorf("Expected error, cards should not be dealt if deck does not have enough cards")
	}
}

func TestShuffle(t *testing.T) {
	deck := NewDeck()
	shuffledDeck := NewDeck()

	shuffledDeck.Shuffle()

	if slices.Compare(deck, shuffledDeck) == 0 {
		t.Errorf("Expected decks to be different, but they were not")
	}
}

func TestAddCard(t *testing.T) {
	deck := NewDeck()

	deck.AddCard("Giga Chad Card")

	if len(deck) != 53 {
		t.Errorf("Expected deck length of 53, but got %v", len(deck))
	}

	if deck[len(deck)-1] != "Giga Chad Card" {
		t.Errorf("Expected last card of Giga Chat Card, but got %v", deck[len(deck)-1])
	}
}
