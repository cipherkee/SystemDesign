package main

import (
	"github.com/cipherkee/SystemDesign/SnakeAndLadderLLD/models"
)

type PlayerQueue struct {
	queue []*models.Player
}

func InitPlayerQueue(players []*models.Player) *PlayerQueue {

	pq := make([]*models.Player, 0)

	for i := 0; i < len(players); i++ {
		pq = append(pq, &models.Player{Name: players[i].Name, Id: players[i].Id})
	}

	return &PlayerQueue{
		queue: pq,
	}
}

func (pq *PlayerQueue) Poll() *models.Player {
	p := pq.queue
	if len(p) > 0 {
		first := p[0]
		if len(p) > 1 {
			p = p[1:]
			pq.queue = p
		} else {
			p = make([]*models.Player, 0)
			pq = &PlayerQueue{
				queue: p,
			}
		}

		return first
	}

	return nil
}

func (pq *PlayerQueue) Push(in *models.Player) {
	p := pq.queue
	p = append(p, in)
	pq.queue = p
}
