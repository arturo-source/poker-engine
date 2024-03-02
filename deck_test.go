package poker_test

import (
	"testing"

	"github.com/arturo-source/poker-engine"
)

func TestGetOneCardFromDeck(t *testing.T) {
	d := poker.NewDeck()

	want := poker.TWOS & poker.FIRST_SUIT
	got := d.GetNextCard()
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestGetTwoCardsFromDeck(t *testing.T) {
	d := poker.NewDeck()

	want := poker.THREES & poker.FIRST_SUIT
	d.GetNextCard()
	got := d.GetNextCard()
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}
