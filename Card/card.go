package card

import "fmt"

// Card represents a playing card as an integer 0 - 51
type Card uint8

var suitMap map[uint8]string = map[uint8]string{
	0: "Hearts",
	1: "Diamonds",
	2: "Clubs",
	3: "Spades",
}

var rankMap map[uint8]string = map[uint8]string{
	1:  "Ace",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
	11: "Jack",
	12: "Queen",
	13: "King",
}

// NewCard returns a pointer to a card
func NewCard(i int) *Card {
	c := Card(i)
	return &c
}

// GetRank returns the rank of the card (1-13)
func (c *Card) GetRank() uint8 {
	return uint8(*c%13) + 1
}

// GetSuit returns the suit of the card
func (c *Card) GetSuit() string {
	k := uint8(*c / 4)
	return suitMap[k]
}

func (c *Card) String() string {
	return fmt.Sprintf("%s of %s", rankMap[c.GetRank()], c.GetSuit())
}
