package card

import (
	"fmt"
)

type Card struct {
	Suit Suit
	Rank Rank
}

func NewCard(suit Suit, rank Rank) *Card {
	return &Card{
		Suit: suit,
		Rank: rank,
	}
}

func (c *Card) String() string {
	return fmt.Sprintf("%s%s", c.Rank.String(), c.Suit.ShortName)
}

func (c *Card) Value() int {
	return int(c.Rank)
}
