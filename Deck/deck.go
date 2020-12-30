package deck

import (
	"blackjack/card"
	"math/rand"
	"time"
)

// Deck is a slice of Cards
type Deck []*card.Card

// NewDeck returns an unshuffled deck of 52 cards
func NewDeck() *Deck {
	d := Deck{}
	for i := 0; i < 52; i++ {
		d = append(d, card.NewCard(i))
	}
	return &d
}

// Shuffle shuffles the deck of cards
func (d *Deck) Shuffle() {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	for i := range *d {
		j := r.Intn(i + 1)
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	}
}

// Pop removes the top card of the deck
func (d *Deck) Pop() *card.Card {
	c := (*d)[len(*d)-1]
	*d = (*d)[:len(*d)-1]
	return c
}
