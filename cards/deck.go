package main

import "fmt"

// Create a new type of "deck"
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamods", "Hearts", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// d here are just the same as "this" in js or "self" in pyton
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// (deck, deck) means that we are return two separte value of the function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

