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

func ExistsSequenceOfThree(c []Card) *Combination {
        if c[0] == c[1] && c[0] == c[2] {
                return &Combination{Card: c[0], Number: 3}
        }

        if c[1] == c[2] && c[1] == c[3] {
                return &Combination{Card: c[1], Number: 3}
        }

        if c[2] == c[3] && c[2] == c[4] {
                return &Combination{Card: c[2], Number: 3}
        }

        if c[0].IsWild() && c[1] == c[2] {
                return &Combination{Card: c[2], Number: 3}
        }    

        if c[1].IsWild() && c[2] == c[3] {
                return &Combination{Card: c[3], Number: 3}
        }

        if c[2].IsWild() && c[3] == c[4] {
                return &Combination{Card: c[4], Number: 3}
        }

        return nil
}

/* func ExistsSequenceOfFour(c []Card) *Combination {
        if c[0] == c[1] && c[1] == c[2] && c[2] == c[3] {
                return &Combination{Card: c[3], Number: 4}
        }

        if c[1] == c[2] && c[2] == c[3] && c[3] == c[4] {
                return &Combination{Card: c[4], Number: 4}
        }

        if c[0].IsWild() && c[1] == c[2] && c[2] == c[3] {
                return &Combination{Card: c[3], Number: 4}
        }

        if c[1].IsWild() && c[2] == c[3] && c[3] == c[4] {
                return &Combination{Card: c[4], Number: 4}
        }

        if c[2].IsWild() && c[0] == c[1] && c[1] == c[3] {
                return &Combination{Card: c[3], Number: 4}
        }

        if c[3].IsWild() && c[2] == c[3] && c[3] == c[4] {
                return &Combination{Card: c[4], Number: 4}
        } 

        return nil
}*/


