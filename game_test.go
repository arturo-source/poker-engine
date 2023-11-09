package poker_test

import (
	"testing"

	"github.com/arturo-source/poker-engine"
)

func TestDealCards(t *testing.T) {
	g := poker.NewGame()
	g.Players = []*poker.Player{poker.NewPlayer("P1")}

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

func TestHighCardTie(t *testing.T) {
	g := poker.NewGame()
	p1 := poker.NewPlayer("P1")
	p2 := poker.NewPlayer("P2")
	g.Players = []*poker.Player{p1, p2}

	c := poker.NewCard
	p1.Hand = c("3h") | c("4h")
	p2.Hand = c("3d") | c("4c")

	g.Board.TableCards = []poker.Card{c("Ad"), c("Kh"), c("Jd"), c("Tc"), c("9h")}

	winners := g.GetWinners()
	if len(winners) != 2 {
		t.Errorf("Expected tie, got %d winners", len(winners))
	}
}

func TestHighCardWins(t *testing.T) {
	g := poker.NewGame()
	p1 := poker.NewPlayer("P1")
	p2 := poker.NewPlayer("P2")
	g.Players = []*poker.Player{p1, p2}

	c := poker.NewCard
	p1.Hand = c("Ah") | c("4h")
	p2.Hand = c("3d") | c("4c")

	g.Board.TableCards = []poker.Card{c("6d"), c("Kh"), c("Jd"), c("Tc"), c("9h")}

	winners := g.GetWinners()
	if len(winners) != 1 {
		t.Errorf("Expected 1 winner, got %d winners", len(winners))
	}

	want := p1
	got := winners[0]
	if want != got {
		t.Errorf("\nWant %v\nGot  %v", want, got)
	}
}

func TestPairTie(t *testing.T) {
	g := poker.NewGame()
	p1 := poker.NewPlayer("P1")
	p2 := poker.NewPlayer("P2")
	g.Players = []*poker.Player{p1, p2}

	c := poker.NewCard
	p1.Hand = c("Ah") | c("Kh")
	p2.Hand = c("Ad") | c("Kc")

	g.Board.TableCards = []poker.Card{c("Kd"), c("3h"), c("4d"), c("6c"), c("7h")}

	winners := g.GetWinners()
	if len(winners) != 2 {
		t.Errorf("Expected tie, got %d winners", len(winners))
	}
}

func TestPairWinsHighCard(t *testing.T) {
	g := poker.NewGame()
	p1 := poker.NewPlayer("P1")
	p2 := poker.NewPlayer("P2")
	g.Players = []*poker.Player{p1, p2}

	c := poker.NewCard
	p1.Hand = c("Ah") | c("Kh")
	p2.Hand = c("5d") | c("5c")

	g.Board.TableCards = []poker.Card{c("Td"), c("3h"), c("9d"), c("6c"), c("7h")}

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

func TestTwoPairTie(t *testing.T) {
	g := poker.NewGame()
	p1 := poker.NewPlayer("P1")
	p2 := poker.NewPlayer("P2")
	g.Players = []*poker.Player{p1, p2}

	c := poker.NewCard
	p1.Hand = c("5h") | c("5s")
	p2.Hand = c("5d") | c("5c")

	g.Board.TableCards = []poker.Card{c("Ks"), c("3s"), c("3d"), c("6c"), c("7h")}

	winners := g.GetWinners()
	if len(winners) != 2 {
		t.Errorf("Expected tie, got %d winners", len(winners))
	}
}

func TestTwoPairWinsHighCard(t *testing.T) {
	g := poker.NewGame()
	p1 := poker.NewPlayer("P1")
	p2 := poker.NewPlayer("P2")
	g.Players = []*poker.Player{p1, p2}

	c := poker.NewCard
	p1.Hand = c("Ah") | c("Kh")
	p2.Hand = c("5d") | c("6d")

	g.Board.TableCards = []poker.Card{c("Td"), c("5h"), c("9d"), c("6c"), c("7h")}

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

func TestThreeOfAKindTie(t *testing.T) {
	g := poker.NewGame()
	p1 := poker.NewPlayer("P1")
	p2 := poker.NewPlayer("P2")
	g.Players = []*poker.Player{p1, p2}

	c := poker.NewCard
	p1.Hand = c("5h") | c("As")
	p2.Hand = c("5d") | c("Ac")

	g.Board.TableCards = []poker.Card{c("Ks"), c("5s"), c("5c"), c("9c"), c("7h")}

	winners := g.GetWinners()
	if len(winners) != 2 {
		t.Errorf("Expected tie, got %d winners", len(winners))
	}
}

func TestThreeOfAKindWinsPair(t *testing.T) {
	g := poker.NewGame()
	p1 := poker.NewPlayer("P1")
	p2 := poker.NewPlayer("P2")
	g.Players = []*poker.Player{p1, p2}

	c := poker.NewCard
	p1.Hand = c("Qd") | c("Qs")
	p2.Hand = c("Ah") | c("Th")

	g.Board.TableCards = []poker.Card{c("Td"), c("Qh"), c("9s"), c("6s"), c("7h")}

	winners := g.GetWinners()
	if len(winners) != 1 {
		t.Errorf("Expected 1 winner, got %d winners", len(winners))
	}

	want := p1
	got := winners[0]
	if want != got {
		t.Errorf("\nWant %v\nGot  %v", want, got)
	}
}

func TestSraightTie(t *testing.T) {
	g := poker.NewGame()
	p1 := poker.NewPlayer("P1")
	p2 := poker.NewPlayer("P2")
	g.Players = []*poker.Player{p1, p2}

	c := poker.NewCard
	p1.Hand = c("Th") | c("Ah")
	p2.Hand = c("Td") | c("Ac")

	g.Board.TableCards = []poker.Card{c("Jc"), c("9h"), c("Ks"), c("8s"), c("7s")}

	winners := g.GetWinners()
	if len(winners) != 2 {
		t.Errorf("Expected tie, got %d winners", len(winners))
	}
}

func TestSraightWinsLowerStraight(t *testing.T) {
	g := poker.NewGame()
	p1 := poker.NewPlayer("P1")
	p2 := poker.NewPlayer("P2")
	g.Players = []*poker.Player{p1, p2}

	c := poker.NewCard
	p1.Hand = c("Th") | c("Ah")
	p2.Hand = c("Td") | c("Qc")

	g.Board.TableCards = []poker.Card{c("Jc"), c("9h"), c("Ks"), c("8s"), c("7s")}

	winners := g.GetWinners()
	if len(winners) != 1 {
		t.Errorf("Expected 1 winner, got %d", len(winners))
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
	g.Players = []*poker.Player{p1, p2}

	c := poker.NewCard
	p1.Hand = c("3h") | c("4h")
	p2.Hand = c("4d") | c("4c")

	g.Board.TableCards = []poker.Card{c("6h"), c("Th"), c("Kh"), c("Ah"), c("5h")}

	winners := g.GetWinners()
	if len(winners) != 2 {
		t.Errorf("Expected tie, got %d winners", len(winners))
	}
}

func TestFlushWinsPair(t *testing.T) {
	g := poker.NewGame()
	p1 := poker.NewPlayer("P1")
	p2 := poker.NewPlayer("P2")
	g.Players = []*poker.Player{p1, p2}

	c := poker.NewCard
	p1.Hand = c("3h") | c("4h")
	p2.Hand = c("4d") | c("4c")

	g.Board.TableCards = []poker.Card{c("6h"), c("Th"), c("Kh"), c("Ad"), c("5d")}

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
