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

func TestPairTie(t *testing.T) {
	g := poker.NewGame()
	p1 := poker.NewPlayer("P1")
	p2 := poker.NewPlayer("P2")
	g.Players = append(g.Players, p1, p2)

	c := poker.NewCard
	p1.Hand = c("Ah") | c("Kh")
	p2.Hand = c("Ad") | c("Kc")

	g.Board.TableCards = append(g.Board.TableCards, c("Kd"), c("3h"), c("4d"), c("6c"), c("7h"))

	winners := g.GetWinners()
	if len(winners) != 2 {
		t.Errorf("Expected tie, got %d winners", len(winners))
	}
}

func TestPairWinsHighCard(t *testing.T) {
	g := poker.NewGame()
	p1 := poker.NewPlayer("P1")
	p2 := poker.NewPlayer("P2")
	g.Players = append(g.Players, p1, p2)

	c := poker.NewCard
	p1.Hand = c("Ah") | c("Kh")
	p2.Hand = c("5d") | c("5c")

	g.Board.TableCards = append(g.Board.TableCards, c("Td"), c("3h"), c("9d"), c("6c"), c("7h"))

	winners := g.GetWinners()
	if len(winners) != 1 {
		t.Errorf("Expected 1 winner, got %d winners", len(winners))
	}

	want := p2
	got := winners[0]
	if want != got {
		t.Errorf("\nWant %v\nGot  %v", want, got)
	}
}

func TestFlushTie(t *testing.T) {
	g := poker.NewGame()
	p1 := poker.NewPlayer("P1")
	p2 := poker.NewPlayer("P2")
	g.Players = append(g.Players, p1, p2)

	c := poker.NewCard
	p1.Hand = c("3h") | c("4h")
	p2.Hand = c("4d") | c("4c")

	g.Board.TableCards = append(g.Board.TableCards, c("6h"), c("Th"), c("Kh"), c("Ah"), c("5h"))

	winners := g.GetWinners()
	if len(winners) != 2 {
		t.Errorf("Expected tie, got %d winners", len(winners))
	}
}

func TestFlushWinsPair(t *testing.T) {
	g := poker.NewGame()
	p1 := poker.NewPlayer("P1")
	p2 := poker.NewPlayer("P2")
	g.Players = append(g.Players, p1, p2)

	c := poker.NewCard
	p1.Hand = c("3h") | c("4h")
	p2.Hand = c("4d") | c("4c")

	g.Board.TableCards = append(g.Board.TableCards, c("6h"), c("Th"), c("Kh"), c("Ad"), c("5d"))

	winners := g.GetWinners()
	if len(winners) != 1 {
		t.Errorf("Expected 1 winner, got %d", len(winners))
	}

	want := p1
	got := winners[0]
	if want != got {
		t.Errorf("\nWant %v\nGot  %v", want, got)
	}
}