package main

import (
	"fmt"

	models "github.com/prasangmisra/socash_card_game/models"
)

func main() {
	var newGame models.Game
	gameResponseFlag := newGame.StartGame()
	if gameResponseFlag {
		fmt.Println("Winner of the game is:", newGame.Winner.Name)
	}
}
