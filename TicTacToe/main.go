package main

import (
	"fmt"
)

func main() {
	var name1, name2 string
	var size int

	fmt.Println("Provide the two players playing the game")
	fmt.Scan(&name1)
	fmt.Scan(&name2)

	fmt.Println("Size of the board now")
	fmt.Scan(&size)

	game := NewGame(size, name1, name2)
	game.PlayGame()
}
