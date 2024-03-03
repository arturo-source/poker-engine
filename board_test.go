package poker_test

import (
	"testing"

	"github.com/arturo-source/poker-engine"
)

func TestBoardState(t *testing.T) {
	d := poker.NewDeck()
	b := poker.NewBoard(d)

	var err error
	var want poker.BoardState
	got := &b.State

	want = poker.PREFLOP
	if want != *got {
		t.Errorf("\nWant %v\nGot  %v", want, *got)
	}

	want = poker.FLOP
	err = b.NextBoardState()
	if err != nil {
		t.Errorf("Got an error flipping cards: %s", err)
	}
	if want != *got {
		t.Errorf("\nWant %v\nGot  %v", want, *got)
	}

	want = poker.TURN
	err = b.NextBoardState()
	if err != nil {
		t.Errorf("Got an error flipping cards: %s", err)
	}
	if want != *got {
		t.Errorf("\nWant %v\nGot  %v", want, *got)
	}

	want = poker.RIVER
	err = b.NextBoardState()
	if err != nil {
		t.Errorf("Got an error flipping cards: %s", err)
	}
	if want != *got {
		t.Errorf("\nWant %v\nGot  %v", want, *got)
	}

	want = poker.SHOWDOWN
	err = b.NextBoardState()
	if err != nil {
		t.Errorf("Got an error flipping cards: %s", err)
	}
	if want != *got {
		t.Errorf("\nWant %v\nGot  %v", want, *got)
	}
}

func TestBoardCardsInFlop(t *testing.T) {
	d := poker.NewDeck()
	b := poker.NewBoard(d)

	err := b.NextBoardState()
	if err != nil {
		t.Errorf("Got an error flipping cards: %s", err)
	}

	want := 3
	got := len(b.TableCards)
	if want != got {
		t.Errorf("\nWant %d\nGot  %d", want, got)
	}

	want = 1
	got = len(b.BurnedCards)
	if want != got {
		t.Errorf("\nWant %d\nGot  %d", want, got)
	}
}
