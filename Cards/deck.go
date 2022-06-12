package main

import "fmt"

// Create a new type of deck which is a slice of strings
type deck []string

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, cardSuits := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuits)
			fmt.Println(cardValue + " of " + cardSuits)
		}
	}
	return cards
}

// d is a type of deck
func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}
