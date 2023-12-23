package main

import "cards/lib/deck"

func main() {
	cardsDeck := deck.NewDeck()
	_, err := cardsDeck.Deal(30)
	if err != nil {
		panic(err)
	}
	cardsDeck.PrintCards()
	cardsDeck.Shuffle()
	cardsDeck.SaveToFile()
	newDeck := deck.NewDeckFromFile()
	newDeck.PrintCards()
}
