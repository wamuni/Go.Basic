package main

import "fmt"

// create a new type of deck
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spade", "Diamonds", "Hearts", "Cub"}
	cardValue := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) shuffle() {

}

func (d deck) deal() {

}

func (d deck) saveToFile() {

}

func (d deck) newDeckFromFile() {

}
