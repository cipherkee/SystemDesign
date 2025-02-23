package models

type Board struct {
	Snakes  map[int]Snake
	Ladders map[int]Ladder

	Size int
}

func NewBoard(snakes, ladders [][]int, size int) *Board {
	snakeMap := map[int]Snake{}
	ladderMap := map[int]Ladder{}

	for i := 0; i < len(snakes); i++ {
		snakeMap[snakes[i][0]] = Snake{start: snakes[i][0], end: snakes[i][1]}
	}
	for i := 0; i < len(ladders); i++ {
		ladderMap[ladders[i][0]] = Ladder{start: ladders[i][0], end: ladders[i][1]}
	}

	return &Board{
		Snakes:  snakeMap,
		Ladders: ladderMap,
		Size:    size,
	}
}

func (b *Board) GetPositionAfterSnakeAndLadder(pos int) int {
	snake, ok := b.Snakes[pos]
	if ok {
		return snake.end // Assuming no snake at the end of a given snake
	}

	ladder, ok := b.Ladders[pos]
	if ok {
		return ladder.end
	}

	return pos
}
