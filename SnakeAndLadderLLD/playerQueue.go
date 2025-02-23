package main

import (
	"github.com/cipherkee/SystemDesign/SnakeAndLadderLLD/models"
)

type PlayerQueue []*models.Player

func InitPlayerQueue(players []*models.Player) PlayerQueue {

	pq := make(PlayerQueue, 0)

	for i := 0; i < len(players); i++ {
		pq = append(pq, &models.Player{Name: players[i].Name, Id: players[i].Id})
	}

	return pq
}

func (p PlayerQueue) Poll() *models.Player {
	if len(p) > 0 {
		first := p[0]
		if len(p) > 1 {
			p = p[1:]
		} else {
			p = make(PlayerQueue, 0)
		}

		return first
	}

	return nil
}

func (p PlayerQueue) Push(in *models.Player) {
	p = append(p, in)
}
