package main

import (
	"errors"
	"fmt"
)

type Mark string

const (
	markempty Mark = ""
	markx     Mark = "x"
	marko     Mark = "o"
)

type Board struct {
	n      int
	marker [][]Mark

	rowCount map[Mark][]int

	colCount map[Mark][]int

	dleft map[Mark]int

	dright map[Mark]int

	totalMarks int
}

func NewBoard(n int) *Board {
	marker := make([][]Mark, n)
	for i := range n {
		marker[i] = make([]Mark, n)
	}

	rowCount := map[Mark][]int{}
	rowCount[markx] = make([]int, n)
	rowCount[marko] = make([]int, n)

	colCount := map[Mark][]int{}
	colCount[markx] = make([]int, n)
	colCount[marko] = make([]int, n)

	dleft := map[Mark]int{}
	dleft[markx] = 0
	dleft[marko] = 0

	dright := map[Mark]int{}
	dright[markx] = 0
	dright[marko] = 0

	return &Board{
		n:        n,
		marker:   marker,
		rowCount: rowCount,
		colCount: colCount,
		dleft:    dleft,
		dright:   dright,
	}
}

func (b *Board) Mark(m Mark, i, j int) (bool, error) {
	if i < 0 || j < 0 || i >= b.n || j >= b.n {
		return false, errors.New("invalid index for marking")
	}

	if b.marker[i][j] != "" {
		return false, errors.New("Already marked index")
	}

	b.totalMarks++
	b.marker[i][j] = m

	b.rowCount[m][i]++

	b.colCount[m][j]++

	if b.rowCount[m][i] == b.n || b.colCount[m][j] == b.n {
		return true, nil
	}

	if i == j {
		b.dleft[m]++
		if b.dleft[m] == b.n {
			return true, nil
		}
	}

	if i+j == b.n-1 {
		b.dright[m]++
		if b.dright[m] == b.n {
			return true, nil
		}
	}
	return false, nil
}

func (b *Board) IsGameOver() bool {
	if b.totalMarks == b.n*b.n {
		return true
	}
	return false
}

func (b *Board) Print() {
	n := b.n
	for i := range n {
		for j := range n {
			if b.marker[i][j] == "" {
				fmt.Print("_")
			} else {
				fmt.Print(b.marker[i][j])
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
