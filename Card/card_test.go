package card

import "testing"

func TestAces(t *testing.T) {
	cards := []Card{Card(0), Card(13), Card(26), Card(39)}

	for _, c := range cards {
		r := c.GetRank()
		if r != 1 {
			t.Errorf("expected 1 got %d", r)
		}
	}
}

func TestKings(t *testing.T) {
	cards := []Card{Card(12), Card(25), Card(38), Card(51)}

	for _, c := range cards {
		r := c.GetRank()
		if r != 13 {
			t.Errorf("expected 13 got %d", r)
		}
	}
}

func TestFive(t *testing.T) {
	cards := []Card{Card(4), Card(17), Card(30), Card(43)}

	for _, c := range cards {
		r := c.GetRank()
		if r != 5 {
			t.Errorf("expected 5 got %d", r)
		}
	}
}

func TestString(t *testing.T) {
	c := Card(0)
	s := c.String()
	expected := "Ace of Hearts"
	if s != expected {
		t.Errorf("expected %s got: %s", expected, s)
	}
}
