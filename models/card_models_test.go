package models

import (
	"testing"
)

func TestGetCardPack(t *testing.T) {
	result := GetCardPack()
	if len(result) != 52 {
		t.Errorf("Expected 56 cards, received %v", len(result))
	}
}

func TestCardSuitsLength(t *testing.T) {
	result := len(getAllCardSuits())
	if result != 4 {
		t.Errorf("Expected 4, received %v", result)
	}
}

func TestCardNames(t *testing.T) {
	result := len(getCardNames())
	if result != 13 {
		t.Errorf("Expected 13, received %v", result)
	}
}

func TestCardNumberPositive(t *testing.T) {
	result := getCardNumber("A")
	if result != 1 {
		t.Errorf("Expected 1, received %v", result)
	}
}

func TestCardNumberNegative(t *testing.T) {
	result := getCardNumber("B")
	if result != -1 {
		t.Errorf("Expected -1, received %v", result)
	}
}

func TestCardWeightPositive(t *testing.T) {
	result := getCardWeight("A")
	if result != 14 {
		t.Errorf("Expected 1, received %v", result)
	}
}

func TestCardWeightNegative(t *testing.T) {
	result := getCardWeight("B")
	if result != -1 {
		t.Errorf("Expected -1, received %v", result)
	}
}
