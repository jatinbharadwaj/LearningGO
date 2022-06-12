package main

import "fmt"

// Create a new type of deck which is a slice of strings
type deck []string

// d is a type of deck
func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}
