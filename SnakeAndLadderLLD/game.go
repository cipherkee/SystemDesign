package main

import (
	"github.com/cipherkee/SystemDesign/SnakeAndLadderLLD/models"
)

type SNLGame struct {
	Board           *models.Board
	PlayerPositions map[*models.Player]int // player to the position
}

func NewSNLGame(players []*models.Player, snake, ladders [][]int, size int) *SNLGame {
	playerPositions := map[*models.Player]int{}
	n := len(players)
	for i := 0; i < n; i++ {
		playerPositions[players[i]] = 0
	}

	board := models.NewBoard(snake, ladders, size)

	return &SNLGame{
		PlayerPositions: playerPositions,
		Board:           board,
	}
}

func (s *SNLGame) MovePosition(p *models.Player, moveCount int) {
	old := s.PlayerPositions[p]
	new := old + moveCount
	new = s.Board.GetPositionAfterSnakeAndLadder(new)
	if new > s.Board.Size {
		return
	}
	s.PlayerPositions[p] = new
}

func (s *SNLGame) IsPlayerWon(p *models.Player) bool {
	pos := s.PlayerPositions[p]
	if pos == s.Board.Size {
		return true
	}

	return false
}
