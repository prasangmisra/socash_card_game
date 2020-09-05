package models

import (
	"testing"
)

func TestCardDistribute(t *testing.T) {
	var g Game
	g.NoOfPlayers = 4
	var players []Player

	var player Player
	players = append(players, player)
	players = append(players, player)
	players = append(players, player)
	players = append(players, player)
	g.Players = players
	result := g.DistributeCardsToPlayersInit(4)
	if !result {
		t.Errorf("Expected something else")
	}
}

func TestCardDistributeNoPlayers(t *testing.T) {
	var g Game
	result := g.DistributeCardsToPlayersInit(4)
	if result {
		t.Errorf("Expected something else")
	}
}

func TestCardDistributeZeroCards(t *testing.T) {
	var g Game
	g.NoOfPlayers = 4
	var players []Player

	var player Player
	players = append(players, player)
	players = append(players, player)
	players = append(players, player)
	players = append(players, player)
	g.Players = players
	result := g.DistributeCardsToPlayersInit(0)
	if result {
		t.Errorf("Expected something else")
	}
}

func TestCardDistributeNegativeCards(t *testing.T) {
	var g Game
	var players []Player
	var player Player
	players = append(players, player)
	players = append(players, player)
	players = append(players, player)
	players = append(players, player)
	g.Players = players
	result := g.DistributeCardsToPlayersInit(-3)
	if result {
		t.Errorf("Expected false since -3 cards cannot be distributed")
	}
}

func TestCardDistributeInsufficientCards(t *testing.T) {
	var g Game
	var players []Player
	var player Player
	players = append(players, player)
	players = append(players, player)
	players = append(players, player)
	players = append(players, player)
	g.Players = players
	result := g.DistributeCardsToPlayersInit(20)
	if result {
		t.Errorf("Expected false since 80 cards will be required")
	}
}
