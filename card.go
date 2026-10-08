package main

import (
	"math/rand"
	"time"
)

type Card int

const (
	Six Card = iota
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Aces
	Wild
)

const CardsTotalWeight = 37
const WildCardIndex = 9

func (c Card) String() string {
	names := [SpinSize]string{"Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Aces", "Wild"}

	return names[int(c)]
}

func (c Card) Price() int {
	prices := [SpinSize]int{6, 7, 8, 9, 10, 15, 20, 25, 30, 1}

	return prices[int(c)]
}

func (c Card) Weight() int {
	weights := [SpinSize]int{4, 4, 4, 4, 4, 4, 4, 4, 4, 1}

	return weights[int(c)]
}

func (c Card) IsWild() bool {
        return int(c) == WildCardIndex
}

func GenerateRandomCardSequence() []Card {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	maxLen := CardsTotalWeight * 2

	cards := make([]Card, 0)
	counter := make(map[Card]int, 0)

	for i := 0; i < maxLen; i++ {
		pos := r.Intn(SpinSize)
		c := Card(pos)

		cnt, ok := counter[c]
		if !ok {
			counter[c] = 1
			cards = append(cards, c)
		} else {
			if cnt == c.Weight() {
				continue
			}
			counter[c]++
			cards = append(cards, c)
		}
	}

	return cards
}
