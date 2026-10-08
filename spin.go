package main

import (
	"math/rand"
	"time"
)

const SpinSize = 9

type Spin struct {
	Cards    []Card
	Position int
}

func NewSpin() *Spin {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	cards := make([]Card, SpinSize)
	rndCards := GenerateRandomCardSequence()

	for ind := 0; ind < SpinSize; ind++ {
		cards[ind] = rndCards[r.Intn(len(rndCards)-1)]
	}

	return &Spin{Cards: cards, Position: r.Intn(SpinSize)}
}

func (s *Spin) Get(shift int) Card {
	if shift < 0 {
		if (shift * -1) >= SpinSize {
			shift = ((shift * -1) % SpinSize) * -1
		}
	} else {
		if shift > SpinSize {
			shift = shift % SpinSize
		}
	}

	pos := s.Position + shift

	if pos < 0 {
		pos = SpinSize + pos
	}

	if pos >= SpinSize {
		pos = pos - SpinSize
	}

	return s.Cards[pos]
}

func (s *Spin) Turn() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s.Position = r.Intn(SpinSize)
}
