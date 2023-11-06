package poker_test

import (
	"testing"

	"github.com/arturo-source/poker-engine"
)

func TestDealCards(t *testing.T) {
	g := poker.NewGame()
	g.Players = append(g.Players, poker.NewPlayer("P1"))

	err := g.DealCards()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	want := 2
	got := g.Players[0].Hand.Ones()
	if want != got {
		t.Errorf("\nWant %d\nGot  %d", want, got)
	}
}
