package deck

import (
	"cards/lib/dir"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const store_filename = "deck.txt"

type Deck []string

func (d *Deck) AddCard(card string) {
	*d = append(*d, card)
}

func (d Deck) PrintCards() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d *Deck) Deal(numCards int) (Deck, error) {
	if numCards > len(*d) {
		error := fmt.Errorf("not enough cards in deck to deal %d cards", numCards)
		return nil, errors.New(error.Error())
	}
	hand, remainingCards := (*d)[:numCards], (*d)[numCards:]
	*d = remainingCards
	return hand, nil
}

func (d Deck) SaveToFile() error {
	deckString := d.ToString()

	os.WriteFile(dir.GetPathname(store_filename), []byte(deckString), 0666)

	fmt.Println("Your deck has been saved successfully")

	return nil
}

func (d Deck) ToString() string {
	return strings.Join(d, "\n")
}

func (d *Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range *d {
		newPosition := r.Intn(len(*d) - 1)
		(*d)[i], (*d)[newPosition] = (*d)[newPosition], (*d)[i]
	}
}

func NewDeck() Deck {
	deck := Deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card := value + " of " + suit
			deck.AddCard(card)
		}
	}

	return deck
}

func NewDeckFromFile() Deck {
	deckBytes, err := os.ReadFile(dir.GetPathname(store_filename))

	if err != nil {
		panic(err)
	}

	return strings.Split(string(deckBytes), "\n")
}
