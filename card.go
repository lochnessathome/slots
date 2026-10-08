package main

import (
	"math/rand/v2"
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

const CardsCount = 10
const CardsTotalWeight = 96
const WildCardIndex = 9

func (c Card) String() string {
	names := [CardsCount]string{"Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Aces", "Wild"}

	return names[int(c)]
}

func (c Card) Price() int {
	prices := [CardsCount]int{6, 7, 8, 9, 10, 15, 20, 25, 30, 1}

	return prices[int(c)]
}

func (c Card) Weight() int {
	weights := [CardsCount]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 1}

	return weights[int(c)]
}

func (c Card) IsWild() bool {
        return int(c) == WildCardIndex
}

func GenerateRandomCardSequence() []Card {
	maxLen := CardsTotalWeight

	cards := make([]Card, 0)
	counter := make(map[Card]int, 0)

	for i := 0; i < maxLen; i++ {
		pos := rand.IntN(CardsCount)
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
