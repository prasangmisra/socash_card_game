package models

import (
	"testing"
)

func TestConstantPlaying(t *testing.T) {
	if PLAYER_STATUS_PLAYING != "playing" {
		t.Errorf("Expected playing, received %v", PLAYER_STATUS_PLAYING)
	}
}
func TestConstantLost(t *testing.T) {
	if PLAYER_STATUS_LOST != "lost" {
		t.Errorf("Expected lost, received %v", PLAYER_STATUS_LOST)
	}
}

func TestConstantTie(t *testing.T) {
	if PLAYER_STATUS_TIE != "tie" {
		t.Errorf("Expected tie, received %v", PLAYER_STATUS_TIE)
	}
}

func TestConstantWin(t *testing.T) {
	if PLAYER_STATUS_WIN != "win" {
		t.Errorf("Expected win, received %v", PLAYER_STATUS_WIN)
	}
}
func TestIsTrailPositive(t *testing.T) {
	var p Player
	p.Cards = append(p.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	p.Cards = append(p.Cards, Card{Suit: "Hearts", Name: "7", Weight: 7, Number: 7})
	p.Cards = append(p.Cards, Card{Suit: "Clubs", Name: "7", Weight: 7, Number: 7})
	result := p.IsTrail()
	if !result {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestIsTrailNegative(t *testing.T) {
	var p Player
	p.Cards = append(p.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	p.Cards = append(p.Cards, Card{Suit: "Hearts", Name: "8", Weight: 8, Number: 8})
	p.Cards = append(p.Cards, Card{Suit: "Clubs", Name: "7", Weight: 7, Number: 7})
	result := p.IsTrail()
	if result {
		t.Errorf("Expected false, got %v", result)
	}
}
func TestIsTrailNegativeLessCards(t *testing.T) {
	var p Player
	p.Cards = append(p.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	p.Cards = append(p.Cards, Card{Suit: "Clubs", Name: "7", Weight: 7, Number: 7})
	result := p.IsTrail()
	if result {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestIsTrailNegativeEmpty(t *testing.T) {
	var p Player
	result := p.IsTrail()
	if result {
		t.Errorf("Expected false, got %v", result)
	}
}
func TestSequencePositive(t *testing.T) {
	var p Player
	p.Cards = append(p.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	p.Cards = append(p.Cards, Card{Suit: "Spades", Name: "8", Weight: 8, Number: 8})
	p.Cards = append(p.Cards, Card{Suit: "Spades", Name: "9", Weight: 9, Number: 9})
	//p.Cards = Shuffle(GetCardPack())[:3]
	result := p.IsSequence()
	if !result {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestSequenceNegative(t *testing.T) {
	var p Player
	p.Cards = append(p.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	p.Cards = append(p.Cards, Card{Suit: "Spades", Name: "8", Weight: 8, Number: 8})
	p.Cards = append(p.Cards, Card{Suit: "Spades", Name: "10", Weight: 10, Number: 10})
	result := p.IsSequence()
	if result {
		t.Errorf("Expected false, got %v", result)
	}
}
func TestSequenceNegativeLessCards(t *testing.T) {
	var p Player
	p.Cards = append(p.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	p.Cards = append(p.Cards, Card{Suit: "Spades", Name: "8", Weight: 8, Number: 8})
	result := p.IsSequence()
	if result {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestSequenceNegativeEmpty(t *testing.T) {
	var p Player
	result := p.IsSequence()
	if result {
		t.Errorf("Expected false, got %v", result)
	}
}
func TestPairPositive(t *testing.T) {
	var p Player
	p.Cards = append(p.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	p.Cards = append(p.Cards, Card{Suit: "Spades", Name: "2", Weight: 2, Number: 2})
	p.Cards = append(p.Cards, Card{Suit: "Hearts", Name: "7", Weight: 7, Number: 7})
	result := p.IsPair()
	if !result {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestPairNegative(t *testing.T) {
	var p Player
	p.Cards = append(p.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	p.Cards = append(p.Cards, Card{Suit: "Spades", Name: "2", Weight: 2, Number: 2})
	p.Cards = append(p.Cards, Card{Suit: "Hearts", Name: "6", Weight: 6, Number: 6})
	result := p.IsPair()
	if result {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestPairNegativeEmpty(t *testing.T) {
	var p Player
	result := p.IsPair()
	if result {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestTopCardPositiveSingle(t *testing.T) {
	var p Player
	p.Cards = append(p.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	p.Cards = append(p.Cards, Card{Suit: "Spades", Name: "2", Weight: 2, Number: 2})
	p.Cards = append(p.Cards, Card{Suit: "Hearts", Name: "8", Weight: 8, Number: 8})
	result := p.GetTopCard()
	if result.Number != 8 {
		t.Errorf("Expected 8, got %v", result.Number)
	}
}
func TestTopCardPositiveMultiple(t *testing.T) {
	var p Player
	p.Cards = append(p.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	p.Cards = append(p.Cards, Card{Suit: "Spades", Name: "2", Weight: 2, Number: 2})
	p.Cards = append(p.Cards, Card{Suit: "Hearts", Name: "7", Weight: 7, Number: 7})
	result := p.GetTopCard()
	if result.Number != 7 {
		t.Errorf("Expected 7, got %v", result.Number)
	}
}
func TestTopCardNegativeNoCard(t *testing.T) {
	var p Player
	result := p.GetTopCard()
	if result.Number != -1 {
		t.Errorf("Expected 0, got %v", result.Number)
	}
}
