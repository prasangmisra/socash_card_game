package models

import (
	"fmt"
	"sort"
)

type Player struct {
	Name   string
	Number int
	Cards  []Card
	Status string
}

const PLAYER_STATUS_PLAYING = "playing"
const PLAYER_STATUS_LOST = "lost"
const PLAYER_STATUS_TIE = "tie"
const PLAYER_STATUS_WIN = "win"

func (p *Player) IsTrail() bool {
	if len(p.Cards) < GAME_PER_PLAYER_INIT {
		return false
	}
	number := p.Cards[0].Number
	fmt.Println(number)
	fmt.Println(len(p.Cards))
	for key := 1; key < len(p.Cards); key++ {
		fmt.Println("Key:", p.Cards[key].Number, ".Number:", number)
		if p.Cards[key].Number != number {
			return false
		}
	}
	return true
}

func (p *Player) IsSequence() bool {
	if len(p.Cards) < GAME_PER_PLAYER_INIT {
		return false
	}
	var numbers []int
	for _, value := range p.Cards {
		numbers = append(numbers, value.Number)
	}
	sort.Ints(numbers)
	for i := 1; i < len(numbers); i++ {
		if numbers[i]-numbers[i-1] != 1 {
			return false
		}
	}
	return true
}

func (p *Player) IsPair() bool {
	if len(p.Cards) < 2 {
		return false
	}
	for outerKey := range p.Cards {
		for innerKey := range p.Cards {
			if outerKey != innerKey && p.Cards[outerKey].Weight == p.Cards[innerKey].Weight {
				fmt.Println("It has a pair")
				return true
			}
		}
	}
	fmt.Println("No Pair found")
	return false
}

func (p *Player) GetTopCard() Card {
	var card Card
	card.Name = ""
	card.Number = -1
	card.Suit = ""
	card.Weight = -1
	highest := -1
	for _, value := range p.Cards {
		if value.Weight > highest {
			highest = value.Weight
			card = value
		}
	}
	return card
}
