package models

import (
	"fmt"
	"math/rand"
	"time"
)

func Something() {
	fmt.Println("Printing")
}

//Card struct, defining card properties
type Card struct {
	Suit   string
	Name   string
	Weight int
	Number int
}

//GetCardPack func will return all the 52 cards
func GetCardPack() []Card {
	var cards []Card
	for _, cardSuitValue := range getAllCardSuits() {
		for _, cardNameValue := range getCardNames() {
			var card Card
			card.Name = cardNameValue
			card.Suit = cardSuitValue
			card.Number = getCardNumber(card.Name)
			card.Weight = getCardWeight(card.Name)
			cards = append(cards, card)
		}
	}
	return cards
}

func Shuffle(vals []Card) []Card {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(vals), func(i, j int) { vals[i], vals[j] = vals[j], vals[i] })
	return vals
}
func getAllCardSuits() []string {
	return []string{"Spades", "Hearts", "Diamonds", "Clubs"}
}

func getCardNames() []string {
	return []string{"A", "K", "Q", "J", "10", "9", "8", "7", "6", "5", "4", "3", "2"}
}

//getCardNumber gives numeric value to all the cards. This follows order K>Q>J>10>...>A
func getCardNumber(input string) int {
	switch input {
	case "A":
		return 1
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 11
	case "10":
		return 10
	case "9":
		return 9
	case "8":
		return 8
	case "7":
		return 7
	case "6":
		return 6
	case "5":
		return 5
	case "4":
		return 4
	case "3":
		return 3
	case "2":
		return 2

	}
	return -1
}

//getCardValue indicated the weight of the card. This follows the order A>K>Q...
func getCardWeight(input string) int {
	switch input {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 11
	case "10":
		return 10
	case "9":
		return 9
	case "8":
		return 8
	case "7":
		return 7
	case "6":
		return 6
	case "5":
		return 5
	case "4":
		return 4
	case "3":
		return 3
	case "2":
		return 2

	}
	return -1
}
