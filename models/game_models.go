package models

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const GAME_PER_PLAYER_INIT = 3

type Game struct {
	Players          []Player
	Winner           Player
	CardsDistributed bool
	CardsLeft        []Card
}

//DistributeCardsToPlayers method will shuffle the pack of cards and distribute
//the fresh set to the players. Number parameter defines how many cards to distribute
func (g *Game) DistributeCardsToPlayersInit(number int) bool {

	//Return false if there are no players or if card number is less than 1
	if len(g.Players) == 0 || number < 1 {
		return false
	}

	//Get All Cards
	g.CardsLeft = Shuffle(GetCardPack())
	//Return false if there are less cards to distribute
	if len(g.CardsLeft) <= (len(g.Players) * number) {
		return false
	}

	//Loop through number param
	//Inside each loop, loop through the players and distribute a card
	count := 0
	for key := 0; key < number; key++ {
		for key1 := range g.Players {
			g.Players[key1].Cards = append(g.Players[key1].Cards, g.CardsLeft[count])
			count++
		}
	}
	//Removing the distributed cards from cards left
	g.CardsLeft = g.CardsLeft[count:]
	g.CardsDistributed = true
	return true
}

//InitGame takes number of players and player details
func (g *Game) InitGamePlayers(in *os.File) bool {
	inReader := bufio.NewReader(in)
	if in == nil {
		in = os.Stdin
		inReader = bufio.NewReader(os.Stdin)
	}
	var numberOfPlayers int
	fmt.Println("Enter the number of players")
	_, err := fmt.Fscanf(in, "%d", &numberOfPlayers)
	if err != nil {
		panic(err)
	}
	//52 is the number of cards available
	if numberOfPlayers > 52 {
		fmt.Println("Cannot play with so many members")
		return false
	}
	var players []Player
	for key := 0; key < numberOfPlayers; key++ {
		var player Player
		fmt.Println("Please enter the name of player ", key+1)
		player.Name, _ = inReader.ReadString('\n')
		if player.Name == "" {
			player.Name = "Player" + strconv.Itoa(key+1)
		}
		fmt.Println(player.Name)
		player.Status = PLAYER_STATUS_PLAYING
		players = append(players, player)
	}
	return true
}

//FindWinnerFromTie gets the single player that has a tie status which is the winner
//If there are more than 1 winner, it will send the first winner
func (g *Game) FindWinnerFromTie() (bool, Player) {
	for key, value := range g.Players {
		if value.Status == PLAYER_STATUS_TIE {
			g.Players[key].Status = PLAYER_STATUS_WIN
			g.Winner = g.Players[key]
			return true, g.Winner
		}
	}
	return false, g.Winner
}

//FindTrailWinner figures if there is any winner from trail
func (g *Game) FindTrailWinner() (bool, Player) {
	trailCount := 0
	for key, value := range g.Players {
		if value.IsTrail() && value.Status == PLAYER_STATUS_PLAYING {
			trailCount++
			g.Players[key].Status = PLAYER_STATUS_TIE
		}
	}
	if trailCount == 1 {
		return g.FindWinnerFromTie()
	}
	return false, g.Winner
}

//FindSequenceWinner figures if there is any winner from sequence
func (g *Game) FindSequenceWinner() (bool, Player) {
	sequenceCount := 0
	for key, value := range g.Players {
		if value.IsSequence() && value.Status == PLAYER_STATUS_PLAYING {
			sequenceCount++
			g.Players[key].Status = PLAYER_STATUS_TIE
		}
	}
	if sequenceCount == 1 {
		return g.FindWinnerFromTie()
	}
	return false, g.Winner
}

//FindPairWinner figures if there is any winner from pair
func (g *Game) FindPairWinner() (bool, Player) {
	pairCount := 0
	for key, value := range g.Players {
		if value.IsPair() && value.Status == PLAYER_STATUS_PLAYING {
			pairCount++
			g.Players[key].Status = PLAYER_STATUS_TIE
		}
	}
	if pairCount == 1 {
		return g.FindWinnerFromTie()
	}
	return false, g.Winner
}

//FindPairWinner figures if there is any winner from pair
func (g *Game) FindTopWinner() (bool, Player) {
	//Get top cards from all players
	var cards []Card
	highestWeight := -1
	topCount := 0
	for _, value := range g.Players {
		cards = append(cards, value.GetTopCard())
		if value.GetTopCard().Weight >= highestWeight {
			highestWeight = value.GetTopCard().Weight
		}
	}
	//Compare the highestWeight of all top cards
	for key, value := range g.Players {
		if value.GetTopCard().Weight == highestWeight {
			g.Players[key].Status = PLAYER_STATUS_TIE
			topCount++
		}
	}
	if topCount == 1 {
		return g.FindWinnerFromTie()
	}
	return false, g.Winner
}

//GetWinner tries to find out if there is any winner after distributing the first set of cards
func (g *Game) GetWinner() (bool, Player) {

	result, winner := g.FindTrailWinner()
	if result {
		return result, winner
	}
	result, winner = g.FindSequenceWinner()
	if result {
		return result, winner
	}
	result, winner = g.FindPairWinner()
	if result {
		return result, winner
	}
	result, winner = g.FindTopWinner()
	if result {
		return result, winner
	}
	return false, g.Winner
}

func (g *Game) StartGame() bool {

	return true
}
