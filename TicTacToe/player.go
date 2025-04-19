package main

type Player struct {
	name string
}

func NewPlayer(name string) *Player {
	return &Player{name: name}
}
