package main

import "fmt"

func main() {

	fmt.Println("Hi There!")

	// Create a new type of deck which is a slice of strings and imported from the package deck
	cards := newDeck()

	// print is imported from the deck package
	cards.print()
}
