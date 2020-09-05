package models

import (
	"io"
	"io/ioutil"
	"os"
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
	result := g.InitGame(in)
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
	result := g.InitGame(in)
	if result {
		t.Errorf("Expected false, got %v", result)
	}
}
