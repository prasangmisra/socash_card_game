package models

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Game struct {
	NoOfPlayers      int
	Players          []Player
	Winner           Player
	CardsDistributed bool
	CardsLeft        []Card
}

func (g Game) GetWinner() Player {
	return g.Winner
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
func (g *Game) InitGame(in *os.File) bool {
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
		players = append(players, player)
	}
	return true
}
