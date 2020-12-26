package card

// Card represents a playing card as an integer 0 - 51
type Card uint8

var suitMap map[uint8]string = map[uint8]string{
	0: "Hearts",
	1: "Diamonds",
	2: "Clubs",
	3: "Spades",
}

// GetRank returns the rank of the card (1-13)
func (c Card) GetRank() uint8 {
	return uint8(c%13) + 1
}

// GetSuit returns the suit of the card
func (c Card) GetSuit() string {
	k := uint8(c / 4)
	return suitMap[k]
}
