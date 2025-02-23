package main

import (
	"fmt"

	"github.com/cipherkee/SystemDesign/SnakeAndLadderLLD/models"
)

type Orchestrator struct {
	Game  *SNLGame
	Dice  *Dice
	Queue PlayerQueue // queue of players by order
}

func InitializeOrchestrator(players []string, boardSize int, snakes, ladders [][]int) *Orchestrator {
	playersModels := makePlayers(players)

	game := NewSNLGame(playersModels, snakes, ladders, boardSize)

	orch := &Orchestrator{
		Dice:  &Dice{},
		Queue: playersModels,
		Game:  game,
	}

	return orch
}

func makePlayers(names []string) []*models.Player {
	n := len(names)
	players := make([]*models.Player, 0)
	for i := 0; i < n; i++ {
		p := &models.Player{
			Name: names[i],
			Id:   fmt.Sprintf("Player%v", i),
		}
		players = append(players, p)
	}

	return players
}

func (o *Orchestrator) RunGame() {

	for {
		currPlayer := o.Queue.Poll()
		roll := o.Dice.Roll()
		fmt.Println(fmt.Sprintf("Player: %v roll: %v", currPlayer.Name, roll))
		o.Game.MovePosition(currPlayer, roll)
		if o.Game.IsPlayerWon(currPlayer) {
			fmt.Println(fmt.Sprintf("Player has won: %s", currPlayer.Name))
			break
		}
		o.Queue.Push(currPlayer)
	}

	fmt.Println("Game is over")
}
