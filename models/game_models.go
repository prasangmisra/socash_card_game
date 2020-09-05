package models

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
