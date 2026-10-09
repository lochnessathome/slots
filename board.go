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

	cm := MatchesFiveWild(c)
	if cm != nil {
		return cm
	}

        cm = MatchesFive(c)
        if cm != nil {
                return cm
        }

        cm = MatchesFourWild(c)
        if cm != nil {
                return cm
        }

        cm = MatchesFour(c)
        if cm != nil {
                return cm
        }

        cm = MatchesThreeWild(c)
        if cm != nil {
                return cm
        }

        cm = MatchesThree(c)
        if cm != nil {
                return cm
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

        cm := MatchesFiveWild(c)
        if cm != nil {
                return cm
        }

        cm = MatchesFive(c)
        if cm != nil {
                return cm
        }

        cm = MatchesFourWild(c)
        if cm != nil {
                return cm
        }

        cm = MatchesFour(c)
        if cm != nil {
                return cm
        }

        cm = MatchesThreeWild(c)
        if cm != nil {
                return cm
        }

        cm = MatchesThree(c)
        if cm != nil {
                return cm
        }

	return nil
}

func (b *Board) SeekRightDiagonalLine(startingShift int) *Combination {
	c := make([]Card, BoardSize)

	for s := BoardSize; s > 0; s-- {
		spin := b.Spins[s - 1]
		card := spin.Get(startingShift + s - 1)
		c[s - 1] = card
	}

        cm := MatchesFiveWild(c)
        if cm != nil {
                return cm
        }

        cm = MatchesFive(c)
        if cm != nil {
                return cm
        }

        cm = MatchesFourWild(c)
        if cm != nil {
                return cm
        }

        cm = MatchesFour(c)
        if cm != nil {
                return cm
        }

        cm = MatchesThreeWild(c)
        if cm != nil {
                return cm
        }

        cm = MatchesThree(c)
        if cm != nil {
                return cm
        }

	return nil
}

