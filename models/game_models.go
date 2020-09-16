package models

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		return false //Either players are 0 or number is less
	}

	//Get All Cards
	g.CardsLeft = Shuffle(GetCardPack())
	//Return false if there are less cards to distribute
	if len(g.CardsLeft) <= (len(g.Players) * number) {
		//fmt.Println("Cannot distribute ", number, "cards to ", len(g.Players), "players")
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

//DistributeFaceOffCards method will distribute single cards to all the players with a tie
func (g *Game) DistributeFaceOffCard() bool {
	for key, value := range g.Players {
		if value.Status == PLAYER_STATUS_PLAYING {
			if len(g.CardsLeft) > 0 {
				//fmt.Println("Adding 1 card for ", g.Players[key].Name)
				g.Players[key].Cards = append([]Card{g.CardsLeft[0]}, g.Players[key].Cards...)
				g.CardsLeft = g.CardsLeft[1:]
			} else {
				//fmt.Println("Cards not found, cards left are ", len(g.CardsLeft))
				return false
			}
		}

	}
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
	//numberOfPlayersString, err := inReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return false
	}
	//numberOfPlayers, _ = strconv.Atoi(numberOfPlayersString)
	//52 is the number of cards available
	if numberOfPlayers > 52 || numberOfPlayers < 2 {
		return false //"Cannot play with so many members")
	}
	var players []Player
	for key := 0; key < numberOfPlayers; key++ {
		var player Player
		fmt.Println("Please enter the name of player ", key+1)
		player.Name, _ = inReader.ReadString('\n')
		player.Name = strings.Trim(player.Name, "\n")
		if player.Name == "" {
			fmt.Println("Name is empty")
			player.Name = "Player" + strconv.Itoa(key+1)
		}
		player.Status = PLAYER_STATUS_PLAYING
		players = append(players, player)
	}
	g.Players = players
	return true
}

//FindAllWinnerFromTie gets the players that have a tie status which is the winner
//If there are more than 1 winner, it will send the first winner
func (g *Game) FindAllWinnerFromTie() (bool, []Player) {
	var result []Player
	result = nil
	resultFlag := false
	for key, value := range g.Players {
		if value.Status == PLAYER_STATUS_TIE {
			g.Players[key].Status = PLAYER_STATUS_WIN
			result = append(result, g.Players[key])
			resultFlag = true
		}
	}

	return resultFlag, result
}

//MarkAllPlaying method marks every player status as playing
func (g *Game) MarkAllPlaying() []Player {
	for key := range g.Players {
		g.Players[key].Status = PLAYER_STATUS_PLAYING
	}
	return g.Players
}

//FindTrailWinner figures if there is any winner from trail
//true if there are 1 or more players with trail
func (g *Game) FindTrailWinner() (bool, []Player) {
	trailCount := 0
	for key, value := range g.Players {
		if value.IsTrail() && value.Status == PLAYER_STATUS_PLAYING {
			trailCount++
			g.Players[key].Status = PLAYER_STATUS_TIE
		}
	}
	//fmt.Println("Trailcount:", trailCount)
	switch trailCount {
	case 0:
		return false, g.MarkAllPlaying()
	case 1:
		_, winner := g.FindAllWinnerFromTie()
		g.Winner = winner[0]
		return true, winner
	default:
		return g.FindAllWinnerFromTie()

	}
}

//FindSequenceWinner figures if there is any winner from sequence
func (g *Game) FindSequenceWinner() (bool, []Player) {
	sequenceCount := 0
	for key, value := range g.Players {
		if value.IsSequence() && value.Status == PLAYER_STATUS_PLAYING {
			sequenceCount++
			g.Players[key].Status = PLAYER_STATUS_TIE
		}
	}
	switch sequenceCount {
	case 0:
		return false, g.MarkAllPlaying()
	case 1:
		_, winner := g.FindAllWinnerFromTie()
		g.Winner = winner[0]
		return true, winner
	default:
		return g.FindAllWinnerFromTie()

	}
}

//FindPairWinner figures if there is any winner from pair
func (g *Game) FindPairWinner() (bool, []Player) {
	pairCount := 0
	for key, value := range g.Players {
		if value.IsPair() && value.Status == PLAYER_STATUS_PLAYING {
			pairCount++
			g.Players[key].Status = PLAYER_STATUS_TIE
		}
	}
	switch pairCount {
	case 0:
		return false, g.MarkAllPlaying()
	case 1:
		_, winner := g.FindAllWinnerFromTie()
		g.Winner = winner[0]
		return true, winner
	default:
		return g.FindAllWinnerFromTie()

	}
}

//FindPairWinner figures if there is any winner from pair
func (g *Game) FindTopWinner() (bool, []Player) {
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
	switch topCount {
	case 0:
		return false, g.MarkAllPlaying()
	case 1:
		_, winner := g.FindAllWinnerFromTie()
		g.Winner = winner[0]
		return true, winner
	default:
		return g.FindAllWinnerFromTie()

	}
}

//GetFirstWinner tries to find out if there is any winner after distributing the first set of cards
func (g *Game) GetFirstWinner() (bool, []Player) {
	result, winner := g.FindTrailWinner()
	if result {
		return result, winner //TrailWinner found
	}
	result, winner = g.FindSequenceWinner()
	if result {
		return result, winner //SequenceWinner found
	}
	result, winner = g.FindPairWinner()
	if result {
		return result, winner //PairWinner found
	}
	result, winner = g.FindTopWinner()
	if result {
		return result, winner //TopWinner found
	}
	//fmt.Println("No Winner found")
	return false, []Player{}
}
func (g *Game) ShowPlayerCards() bool {
	for _, value := range g.Players {
		for _, value1 := range value.Cards {
			fmt.Println("\t", value1.Suit, value1.Name)
		}
		fmt.Println()
	}
	return true
}

func (g *Game) UpdateStatusOfPlayers() bool {
	for key, value := range g.Players {
		//fmt.Println(value.Name, "-", value.Status)
		switch value.Status {
		case PLAYER_STATUS_PLAYING:
			g.Players[key].UpdateStatus(PLAYER_STATUS_LOST)
			break
		case PLAYER_STATUS_WIN:
			g.Players[key].UpdateStatus(PLAYER_STATUS_PLAYING)
			break
		default:
			break
		}

	}
	return true
}
func (g *Game) PlayFaceOff() bool {
	for len(g.CardsLeft) > 0 {
		cardFlag := g.DistributeFaceOffCard()
		if !cardFlag {
			return false
		}
		//Check First card of all the playing players
		heighestWeight := -1
		topCount := 0
		for _, value := range g.Players {
			if value.Status == PLAYER_STATUS_PLAYING && value.Cards[0].Weight >= heighestWeight {
				heighestWeight = value.Cards[0].Weight
			}
		}
		//fmt.Println(heighestWeight, "is the highestWeight")
		for key, value := range g.Players {
			if value.Status == PLAYER_STATUS_PLAYING && value.Cards[0].Weight == heighestWeight {
				topCount++
				g.Players[key].Status = PLAYER_STATUS_TIE
			} else {
				g.Players[key].Status = PLAYER_STATUS_LOST
			}
		}
		//fmt.Println("Highestweight found in ", topCount, "users")
		if topCount == 1 {
			//fmt.Println("Only one found")
			for key, value := range g.Players {
				if value.Status == PLAYER_STATUS_TIE {
					g.Players[key].Status = PLAYER_STATUS_WIN
					g.Winner = g.Players[key]
					return true
				}
			}

		}
		//If there is 1 player with highest weight, that is the winner
		//If there are more than 1 players with highest weight card, only mark the other players as lost, the highest players will be playing now

	}
	return false
}
func (g *Game) StartGame() bool {
	stepResponse := g.InitGamePlayers(nil)
	if stepResponse {
		stepResponse = g.DistributeCardsToPlayersInit(GAME_PER_PLAYER_INIT)
		if stepResponse {
			//g.ShowPlayerCards()
			fmt.Println("Cards are distributed")
			stepResponse, players := g.GetFirstWinner()
			if stepResponse && len(players) > 0 {
				//fmt.Println("All Good. Total winners are:", len(players))
				if len(players) == 1 {

					fmt.Println("Game over, winner is ", g.Winner.Name)
					return true
				} else {
					g.UpdateStatusOfPlayers()
					fmt.Println("Will have to send some more cards")
					return g.PlayFaceOff()
				}
			} else {
				//fmt.Println("stepResponse:", stepResponse)
				//fmt.Println("Players:", players)
				return false
			}
		} else {
			//fmt.Println("Something happened at DistributeCardsToPlayersInit")
		}
		return false
	} else {
		//fmt.Println("Something happened at InitGamePlayers")
	}
	return false
}
