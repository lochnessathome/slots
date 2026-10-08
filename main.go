package main

import (
	"fmt"
)

const (
	attempts = 1000
	rate     = 0.03
)

func main() {
	board := NewBoard()

	counter := 0
	wins := 0

	for a := 0; a < attempts; a++ {
		counter++

		if !winRateFits(counter, wins) {
			for {
				combination := seek(board)
				if combination != nil {
					board.Turn()
				} else {
					break
				}
			}
		} else {
			combination := seek(board)
			if combination != nil {
				fmt.Println(board.Sprintf())
				combination.Print()
				wins++
			}
		}

		board.Turn()
	}

	fmt.Println(wins, "of", attempts, "=", (100.0 * float32(wins) / float32(attempts)), "%")
}

func seek(board *Board) *Combination {
	for r := -2; r <= 2; r++ {
		combination := board.SeekHorizontalLine(r)

		if combination != nil {
			return combination
		}
	}

	return nil
}

func winRateFits(counter, wins int) bool {
	if (float32(wins) / float32(counter)) > rate {
		return false
	}

	return true
}
