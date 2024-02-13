package main

import "fmt"

func main() {
	cards := newDeck()

	inhand, indeck := deal(cards, 5)

	inhand.print()
	indeck.print()

	var newDeck deck

	newDeck, err := newDeck.newDeckFromFile("hand.txt")
	if err != nil {
		fmt.Println("Error occured when loading file", err)
	}
	newDeck.print()

	if inhand.saveToFile("hand.txt") != nil {
		fmt.Println("Saving hand.txt occured error...")
	}

	if indeck.saveToFile("deck.txt") != nil {
		fmt.Println("Saving deck.txt occured error...")
	}
}
