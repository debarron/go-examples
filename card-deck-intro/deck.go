package main

import "fmt"

type deck []string

func newDeck() deck {
	cardSuits := []string{"Spades", "Hearts", "Cloves", "Dimonds"}
	cardNumbers := []string{
		"A", "J", "Q", "K",
		"2", "3", "4", "5", "6", "7",
		"8", "9", "10"}

	var cardDeck deck
	for _, suit := range cardSuits {
		for _, number := range cardNumbers {
			cardDeck = append(cardDeck, number+" of "+suit)
		}
	}

	return cardDeck
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[handSize:], d[:handSize]
}
