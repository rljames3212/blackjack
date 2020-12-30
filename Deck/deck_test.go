package deck

import "testing"

func TestShuffle(t *testing.T) {
	d := NewDeck()
	d.Shuffle()
	unshuffled := true
	for i, c := range *d {
		if int(*c) != i {
			unshuffled = false
		}
	}
	if unshuffled {
		t.Error("Deck is not shuffled")
	}
}

func TestPop(t *testing.T) {
	d := NewDeck()
	d.Shuffle()

	topCard := (*d)[len(*d)-1]
	originalLength := len(*d)

	c := d.Pop()

	expectedLength := originalLength - 1

	if c != topCard {
		t.Errorf("expected %s as top card, got %s as top card", topCard.String(), c.String())
	}
	if len(*d) != expectedLength {
		t.Errorf("expected %d as length, got %d as length", expectedLength, len(*d))
	}

}
