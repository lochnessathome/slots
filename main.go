package main

import (
	"fmt"
)

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
