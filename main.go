package main

import (
	"fmt"
)

type Combination struct {
	Card   Card
	Number int
}

func (c *Combination) Print() {
	fmt.Println("card:", c.Card.String(), "number:", c.Number, "price:", c.Card.Price(), "total:", c.Card.Price()*c.Number)
}

func main() {
	board := NewBoard()

	fmt.Printf(board.Sprintf())

	for r := -2; r <= 2; r++ {
		combination := board.SeekHorizontalLine(r)
		if combination != nil {
			fmt.Println("line:", r)
			combination.Print()
		}
	}

	diaCombination := board.SeekLeftDiagonalLine(-2)
	if diaCombination != nil {
		fmt.Println("left diagonal")
		diaCombination.Print()
	}

}
