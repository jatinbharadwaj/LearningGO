package main

import "fmt"

func main() {

	var card string = "Ace of Spades"
	NewCard := "Five of Diamonds" // Same As Above
	NewCardWithFunction := newCardWithFunction()
	fmt.Println(card)
	fmt.Println(NewCard)
	fmt.Println(NewCardWithFunction)

}
func newCardWithFunction() string {
	// string literal implies the return type is String
	return "Ace of Diamonds"
}
