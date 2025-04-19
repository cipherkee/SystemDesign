package main

import (
	"fmt"
)

type Game struct {
	board         *Board
	p1            *Player
	p2            *Player
	playerMarkMap map[*Player]Mark
}

func NewGame(n int, name1, name2 string) *Game {
	p1 := NewPlayer(name1)
	p2 := NewPlayer(name2)
	return &Game{
		board: NewBoard(n),
		p1:    p1,
		p2:    p2,
		playerMarkMap: map[*Player]Mark{
			p1: markx,
			p2: marko,
		},
	}
}

func (g *Game) PlayGame() {
	p := g.p1
	anywon := false

	fmt.Print(g.board.n)

	for !g.board.IsGameOver() {
		var i, j int
		fmt.Printf("Player %v, provide input for i and j \n", p.name)
		fmt.Scan(&i)
		fmt.Scan(&j)
		mark := g.playerMarkMap[p]
		won, err := g.board.Mark(mark, i, j)
		if err != nil {
			fmt.Println("Incorrect values for player, play again", err.Error())
			continue
		}
		if won {
			anywon = true
			fmt.Printf("Player %v has won the game. Congradulations!!!\n", p.name)
			g.board.Print()
			break
		}
		if p == g.p1 {
			p = g.p2
		} else {
			p = g.p1
		}
		g.board.Print()
	}
	if !anywon {
		fmt.Println("Match draw!!!")
	}
}
