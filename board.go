package main

import (
	"fmt"
)

const BoardSize = 5

type Board struct {
	Spins [5]*Spin
}

func NewBoard() *Board {
	board := &Board{}

	for i := 0; i < BoardSize; i++ {
		board.Spins[i] = NewSpin()
	}

	return board
}

func (b *Board) Turn() {
	for i := 0; i < BoardSize; i++ {
		b.Spins[i].Turn()
	}
}

func (b *Board) Sprintf() string {
	var out string

	for r := -2; r <= 2; r++ {
		for c := 0; c < BoardSize; c++ {
			spin := b.Spins[c]
			card := spin.Get(r)

			out = fmt.Sprintf("%s%s\t", out, card.String())

			if c == (BoardSize - 1) {
				out = fmt.Sprintf("%s\n", out)
			}
		}
	}

	return out
}

func (b *Board) SeekHorizontalLine(shift int) *Combination {
	c := make([]Card, BoardSize)

	for s := 0; s < BoardSize; s++ {
		spin := b.Spins[s]
		card := spin.Get(shift)
		c[s] = card
	}

	if c[0] == c[1] && c[0] == c[2] && c[0] == c[3] && c[0] == c[4] {
		return &Combination{Card: c[0], Number: 5}
	}

	if c[0] == c[1] && c[0] == c[2] && c[0] == c[3] {
		return &Combination{Card: c[0], Number: 4}
	}

	if c[1] == c[2] && c[1] == c[3] && c[1] == c[4] {
		return &Combination{Card: c[1], Number: 4}
	}

	ct := ExistsSequenceOfThree(c)
	if ct != nil {
		return ct
	}

	return nil
}

func (b *Board) SeekLeftDiagonalLine(startingShift int) *Combination {
	c := make([]Card, BoardSize)

	for s := 0; s < BoardSize; s++ {
		spin := b.Spins[s]
		card := spin.Get(startingShift + s)
		c[s] = card
	}

	if c[0] == c[1] && c[0] == c[2] && c[0] == c[3] && c[0] == c[4] {
		return &Combination{Card: c[0], Number: 5}
	}

	if c[0] == c[1] && c[0] == c[2] && c[0] == c[3] {
		return &Combination{Card: c[0], Number: 4}
	}

	if c[1] == c[2] && c[1] == c[3] && c[1] == c[4] {
		return &Combination{Card: c[1], Number: 4}
	}

	if c[0] == c[1] && c[0] == c[2] {
		return &Combination{Card: c[0], Number: 3}
	}

	if c[1] == c[2] && c[1] == c[3] {
		return &Combination{Card: c[1], Number: 3}
	}

	if c[2] == c[3] && c[2] == c[4] {
		return &Combination{Card: c[2], Number: 3}
	}

	return nil
}
