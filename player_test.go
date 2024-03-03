package poker_test

import (
	"testing"

	"github.com/arturo-source/poker-engine"
)

func TestAddOneCard(t *testing.T) {
	p := poker.NewPlayer("")
	card := poker.NewCard("As")

	want := card
	got := &p.Hand

	err := p.AddCard(card)
	if err != nil {
		t.Errorf("Got error adding a card to a player hand: %s", err)
	}

	if want != *got {
		t.Errorf("\nWant %s\nGot  %s", want, *got)
	}
}

func TestAddTwoCards(t *testing.T) {
	p := poker.NewPlayer("")
	card1 := poker.NewCard("As")
	card2 := poker.NewCard("Kd")

	want := card1 | card2
	got := &p.Hand

	var err error
	err = p.AddCard(card1)
	if err != nil {
		t.Errorf("Got error adding a card to a player hand: %s", err)
	}
	err = p.AddCard(card2)
	if err != nil {
		t.Errorf("Got error adding a card to a player hand: %s", err)
	}

	if want != *got {
		t.Errorf("\nWant %s\nGot  %s", want, *got)
	}
}

func TestErrorAddingMoreThanMaxCards(t *testing.T) {
	p := poker.NewPlayer("")
	card1 := poker.NewCard("As")
	card2 := poker.NewCard("Kd")
	card3 := poker.NewCard("3c")

	var err error
	err = p.AddCard(card1)
	if err != nil {
		t.Errorf("Got error adding a card to a player hand: %s", err)
	}
	err = p.AddCard(card2)
	if err != nil {
		t.Errorf("Got error adding a card to a player hand: %s", err)
	}

	err = p.AddCard(card3)
	if err == nil {
		t.Errorf("Wanted an error. Got nil.")
	}
}
