package models

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestCardDistributeNoPlayers(t *testing.T) {
	var g Game
	result := g.DistributeCardsToPlayersInit(4)
	if result {
		t.Errorf("Expected false since there are 0 players, received %v", result)
	}
}
func TestCardDistributeZeroCards(t *testing.T) {
	var g Game
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
func TestCardDistribute(t *testing.T) {
	var g Game
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

func TestInitGame(t *testing.T) {
	var g Game
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()
	_, err = io.WriteString(in, "4\n"+"a\nb\nc\nd")
	if err != nil {
		t.Fatal(err)
	}
	_, err = in.Seek(0, os.SEEK_SET)
	if err != nil {
		t.Fatal(err)
	}
	result := g.InitGamePlayers(in)
	if !result {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestInitGameNegativeMorePlayers(t *testing.T) {
	var g Game
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()
	_, err = io.WriteString(in, "53\n"+"Prasang Misra\nb\nc\nd")
	if err != nil {
		t.Fatal(err)
	}
	_, err = in.Seek(0, os.SEEK_SET)
	if err != nil {
		t.Fatal(err)
	}
	result := g.InitGamePlayers(in)
	if result {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestInitGameNegativeLessPlayers(t *testing.T) {
	var g Game
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()
	_, err = io.WriteString(in, "1\n"+"Prasang Misra\nb\nc\nd")
	if err != nil {
		t.Fatal(err)
	}
	_, err = in.Seek(0, os.SEEK_SET)
	if err != nil {
		t.Fatal(err)
	}
	result := g.InitGamePlayers(in)
	if result {
		t.Errorf("Expected false, got %v", result)
	}
}
func TestFindWinnerFromTiePositive(t *testing.T) {
	var g Game
	var player Player
	player.Status = PLAYER_STATUS_TIE
	player.Name = "A"
	g.Players = append(g.Players, player)
	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "B"
	g.Players = append(g.Players, player)
	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "C"
	g.Players = append(g.Players, player)

	resultFlag, result := g.FindAllWinnerFromTie()
	if !resultFlag || result[0].Name != "A" {
		t.Errorf("Expected A as the winner, got %v", result[0].Name)
	}
}

func TestFindWinnerFromTiePositiveMultiple(t *testing.T) {
	var g Game
	var player Player
	player.Status = PLAYER_STATUS_TIE
	player.Name = "A"
	g.Players = append(g.Players, player)
	player.Status = PLAYER_STATUS_TIE
	player.Name = "B"
	g.Players = append(g.Players, player)
	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "C"
	g.Players = append(g.Players, player)

	resultFlag, result := g.FindAllWinnerFromTie()
	fmt.Println(resultFlag)
	if !resultFlag || result[0].Name != "A" {
		t.Errorf("Expected A as the winner, got %v", result[0].Name)
	}
}

func TestFindWinnerFromTieNegativeNoWinner(t *testing.T) {
	var g Game
	var player Player
	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "A"
	g.Players = append(g.Players, player)
	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "B"
	g.Players = append(g.Players, player)
	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "C"
	g.Players = append(g.Players, player)

	resultFlag, result := g.FindAllWinnerFromTie()
	if resultFlag || len(result) != 0 {
		t.Errorf("Expected Blank as the winner, got %v results", len(result))
	}
}

func TestFindWinnerFromTieNegativeEmpty(t *testing.T) {
	var g Game

	resultFlag, result := g.FindAllWinnerFromTie()
	if resultFlag || len(result) > 0 {
		t.Errorf("Expected Blank as the winner, got %v %v", len(result), resultFlag)
	}
}

func TestFindTrailWinnerPositive(t *testing.T) {
	var g Game
	var player Player

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "A"
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "7", Weight: 7, Number: 7})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "B"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "8", Weight: 8, Number: 8})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "9", Weight: 9, Number: 9})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "C"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "9", Weight: 9, Number: 9})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "8", Weight: 8, Number: 8})
	g.Players = append(g.Players, player)

	resultFlag, result := g.FindTrailWinner()
	if !resultFlag || result[0].Name != "A" {
		t.Errorf("Expected none, got %v and flag %v", result[0].Name, resultFlag)
	}
}

func TestFindTrailWinnerNegative(t *testing.T) {
	var g Game
	var player Player

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "A"
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "3", Weight: 3, Number: 3})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "7", Weight: 7, Number: 7})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "B"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "8", Weight: 8, Number: 8})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "9", Weight: 9, Number: 9})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "C"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "9", Weight: 9, Number: 9})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "8", Weight: 8, Number: 8})
	g.Players = append(g.Players, player)

	resultFlag, result := g.FindTrailWinner()
	if resultFlag || len(result) != 3 {
		t.Errorf("Expected 3 and false, got %v and flag %v", len(result), resultFlag)
	}
}

func TestFindTrailWinnerNegativeMultipleTrail(t *testing.T) {
	var g Game
	var player Player

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "A"
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "7", Weight: 7, Number: 7})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "B"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "7", Weight: 7, Number: 7})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "C"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "9", Weight: 9, Number: 9})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "8", Weight: 8, Number: 8})
	g.Players = append(g.Players, player)

	resultFlag, result := g.FindTrailWinner()
	if !resultFlag || len(result) != 2 {
		t.Errorf("Expected none, got %v and flag %v", len(result), resultFlag)
	}
}

func TestFindTrailWinnerNegativeNoPlayers(t *testing.T) {
	var g Game
	resultFlag, result := g.FindTrailWinner()
	if resultFlag || len(result) != 0 {
		t.Errorf("Expected none, got %v and flag %v", result[0].Name, resultFlag)
	}
}

func TestFindSequenceWinnerPositive(t *testing.T) {
	var g Game
	var player Player

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "A"
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "7", Weight: 7, Number: 7})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "B"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "8", Weight: 8, Number: 8})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "9", Weight: 9, Number: 9})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "C"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "9", Weight: 9, Number: 9})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "3", Weight: 3, Number: 3})
	g.Players = append(g.Players, player)

	resultFlag, result := g.FindSequenceWinner()
	if !resultFlag || result[0].Name != "B" {
		t.Errorf("Expected none, got %v and flag %v", result[0].Name, resultFlag)
	}
}

func TestFindSequenceWinnerNegative(t *testing.T) {
	var g Game
	var player Player

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "A"
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "3", Weight: 3, Number: 3})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "7", Weight: 7, Number: 7})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "B"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "2", Weight: 2, Number: 2})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "4", Weight: 4, Number: 4})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "C"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "9", Weight: 9, Number: 9})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "2", Weight: 2, Number: 2})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "5", Weight: 5, Number: 5})
	g.Players = append(g.Players, player)

	resultFlag, result := g.FindSequenceWinner()
	if resultFlag || len(result) != 3 {
		t.Errorf("Expected 3 and false, got %v and flag %v", len(result), resultFlag)
	}
}

func TestFindSequenceWinnerMultipleSequence(t *testing.T) {
	var g Game
	var player Player

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "A"
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "8", Weight: 8, Number: 8})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "9", Weight: 9, Number: 9})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "B"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "8", Weight: 8, Number: 8})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "9", Weight: 9, Number: 9})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "C"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "9", Weight: 9, Number: 9})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "8", Weight: 8, Number: 8})
	g.Players = append(g.Players, player)

	resultFlag, result := g.FindSequenceWinner()
	if !resultFlag || len(result) != 3 {
		t.Errorf("Expected 3 and true, got %v and flag %v", len(result), resultFlag)
	}
}

func TestFindSequenceWinnerNegativeNoPlayers(t *testing.T) {
	var g Game
	resultFlag, result := g.FindSequenceWinner()
	if resultFlag || len(result) != 0 {
		t.Errorf("Expected none, got %v and flag %v", result[0].Name, resultFlag)
	}
}

func TestFindPairWinnerPositive(t *testing.T) {
	var g Game
	var player Player

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "A"
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "7", Weight: 7, Number: 7})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "B"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "8", Weight: 8, Number: 8})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "9", Weight: 9, Number: 9})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "C"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "9", Weight: 9, Number: 9})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "3", Weight: 3, Number: 3})
	g.Players = append(g.Players, player)

	resultFlag, result := g.FindPairWinner()
	if !resultFlag || result[0].Name != "A" {
		t.Errorf("Expected none, got %v and flag %v", result[0].Name, resultFlag)
	}
}

func TestFindPairWinnerNegative(t *testing.T) {
	var g Game
	var player Player

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "A"
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "3", Weight: 3, Number: 3})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "6", Weight: 6, Number: 6})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "B"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "2", Weight: 2, Number: 2})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "4", Weight: 4, Number: 4})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "C"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "9", Weight: 9, Number: 9})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "2", Weight: 2, Number: 2})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "5", Weight: 5, Number: 5})
	g.Players = append(g.Players, player)

	resultFlag, result := g.FindPairWinner()
	if resultFlag || len(result) != 3 {
		t.Errorf("Expected none, got %v and flag %v", len(result), resultFlag)
	}
}

func TestFindPairWinnerMultiplePairs(t *testing.T) {
	var g Game
	var player Player

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "A"
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "9", Weight: 9, Number: 9})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "B"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "8", Weight: 8, Number: 8})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "8", Weight: 8, Number: 8})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "C"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "9", Weight: 9, Number: 9})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "8", Weight: 8, Number: 8})
	g.Players = append(g.Players, player)

	resultFlag, result := g.FindPairWinner()
	if !resultFlag || len(result) != 2 {
		t.Errorf("Expected none, got %v and flag %v", result[0].Name, resultFlag)
	}
}

func TestFindPairWinnerNegativeNoPlayers(t *testing.T) {
	var g Game
	resultFlag, result := g.FindPairWinner()
	if resultFlag || len(result) != 0 {
		t.Errorf("Expected none, got %v and flag %v", len(result), resultFlag)
	}
}

func TestFindTopWinnerPositive(t *testing.T) {
	var g Game
	var player Player

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "A"
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "7", Weight: 7, Number: 7})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "B"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "8", Weight: 8, Number: 8})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "9", Weight: 9, Number: 9})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "C"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "4", Weight: 4, Number: 4})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "3", Weight: 3, Number: 3})
	g.Players = append(g.Players, player)

	resultFlag, result := g.FindTopWinner()
	if !resultFlag || result[0].Name != "B" {
		t.Errorf("Expected B, got %v and flag %v", result[0].Name, resultFlag)
	}
}

func TestFindTopWinnerMultiple(t *testing.T) {
	var g Game
	var player Player

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "A"
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "3", Weight: 3, Number: 3})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "6", Weight: 6, Number: 6})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "B"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "2", Weight: 2, Number: 2})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "4", Weight: 4, Number: 4})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "C"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "3", Weight: 3, Number: 3})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "2", Weight: 2, Number: 2})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "5", Weight: 5, Number: 5})
	g.Players = append(g.Players, player)

	resultFlag, result := g.FindTopWinner()
	if !resultFlag || len(result) != 2 {
		t.Errorf("Expected 2, got %v and flag %v", len(result), resultFlag)
	}
}

func TestFindTopWinnerNegativeNoPlayers(t *testing.T) {
	var g Game
	resultFlag, result := g.FindTopWinner()
	if resultFlag || len(result) != 0 {
		t.Errorf("Expected none, got %v and flag %v", len(result), resultFlag)
	}
}

func TestDistributeFaceOffCardsPositive(t *testing.T) {
	var g Game
	var player Player

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "A"
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "9", Weight: 9, Number: 9})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "B"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "8", Weight: 8, Number: 8})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "9", Weight: 9, Number: 9})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "8", Weight: 8, Number: 8})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "C"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "4", Weight: 4, Number: 4})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "8", Weight: 8, Number: 8})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "7", Weight: 7, Number: 7})
	g.Players = append(g.Players, player)
	g.CardsLeft = GetCardPack()
	for key := 0; key < len(g.CardsLeft); key++ {
		value := g.CardsLeft[key]
		if (value.Suit == "Spades" && (value.Name == "7" || value.Name == "8" || value.Name == "4")) ||
			(value.Suit == "Hearts" && (value.Name == "7" || value.Name == "9" || value.Name == "8")) ||
			(value.Suit == "Clubs" && (value.Name == "9" || value.Name == "8" || value.Name == "7")) {
			g.CardsLeft = append(g.CardsLeft[:key-1], g.CardsLeft[key:]...)
			key--
		}
	}
	result, _ := g.GetFirstWinner()
	if !result {
		finalResult := g.DistributeFaceOffCard()
		if finalResult {
			if len(g.Players[0].Cards) != 4 {
				t.Errorf("Expected 5 cards, received %v for %v", len(g.Players[0].Cards), g.Players[0].Name)
			}
		}
	}
}

func TestDistributeFaceOffCardsNegativeNoCardsLeft(t *testing.T) {
	var g Game
	var player Player

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "A"
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "9", Weight: 9, Number: 9})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "B"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "8", Weight: 8, Number: 8})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "9", Weight: 9, Number: 9})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "8", Weight: 8, Number: 8})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "C"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "4", Weight: 4, Number: 4})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "8", Weight: 8, Number: 8})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "7", Weight: 7, Number: 7})
	g.Players = append(g.Players, player)
	result, _ := g.GetFirstWinner()
	if !result {
		finalResult := g.DistributeFaceOffCard()
		if finalResult {
			t.Errorf("0 Cards left, still giving true")
		}
	}
}

func TestShowPlayerCards(t *testing.T) {
	var g Game
	var player Player

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "A"
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "9", Weight: 9, Number: 9})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "B"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "8", Weight: 8, Number: 8})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "9", Weight: 9, Number: 9})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "8", Weight: 8, Number: 8})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "C"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "4", Weight: 4, Number: 4})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "8", Weight: 8, Number: 8})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "7", Weight: 7, Number: 7})
	g.Players = append(g.Players, player)
	result := g.ShowPlayerCards()
	if !result {
		t.Errorf("Expected true, received %v", result)
	}
}
func TestUpdateStatusOfPlayersPlaying(t *testing.T) {
	var g Game
	var player Player

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "A"
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "7", Weight: 7, Number: 7})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "9", Weight: 9, Number: 9})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_PLAYING
	player.Name = "B"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "8", Weight: 8, Number: 8})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "9", Weight: 9, Number: 9})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "8", Weight: 8, Number: 8})
	g.Players = append(g.Players, player)

	player.Status = PLAYER_STATUS_WIN
	player.Name = "C"
	player.Cards = nil
	player.Cards = append(player.Cards, Card{Suit: "Spades", Name: "4", Weight: 4, Number: 4})
	player.Cards = append(player.Cards, Card{Suit: "Hearts", Name: "8", Weight: 8, Number: 8})
	player.Cards = append(player.Cards, Card{Suit: "Clubs", Name: "7", Weight: 7, Number: 7})
	g.Players = append(g.Players, player)
	result := g.UpdateStatusOfPlayers()
	if !result {
		t.Errorf("Expected true, received %v", result)
	}
}
