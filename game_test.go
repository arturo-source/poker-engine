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
	got := g.Players[0].Hand.Count()
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

	g.Board.TableCards = []poker.Cards{c("Ad"), c("Kh"), c("Jd"), c("Tc"), c("9h")}
	tableCards := poker.JoinCards(g.Board.TableCards...)

	winners := poker.GetWinners(tableCards, g.Players)
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

	g.Board.TableCards = []poker.Cards{c("6d"), c("Kh"), c("Jd"), c("Tc"), c("9h")}
	tableCards := poker.JoinCards(g.Board.TableCards...)

	winners := poker.GetWinners(tableCards, g.Players)
	if len(winners) != 1 {
		t.Errorf("Expected 1 winner, got %d winners", len(winners))
	}

	want := p1
	got := winners[0]
	if want != got.Player {
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

	g.Board.TableCards = []poker.Cards{c("Kd"), c("3h"), c("4d"), c("6c"), c("7h")}
	tableCards := poker.JoinCards(g.Board.TableCards...)

	winners := poker.GetWinners(tableCards, g.Players)
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

	g.Board.TableCards = []poker.Cards{c("Td"), c("3h"), c("9d"), c("6c"), c("7h")}
	tableCards := poker.JoinCards(g.Board.TableCards...)

	winners := poker.GetWinners(tableCards, g.Players)
	if len(winners) != 1 {
		t.Errorf("Expected 1 winner, got %d winners", len(winners))
	}

	want := p2
	got := winners[0]
	if want != got.Player {
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

	g.Board.TableCards = []poker.Cards{c("Ks"), c("3s"), c("3d"), c("6c"), c("7h")}
	tableCards := poker.JoinCards(g.Board.TableCards...)

	winners := poker.GetWinners(tableCards, g.Players)
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

	g.Board.TableCards = []poker.Cards{c("Td"), c("5h"), c("9d"), c("6c"), c("7h")}
	tableCards := poker.JoinCards(g.Board.TableCards...)

	winners := poker.GetWinners(tableCards, g.Players)
	if len(winners) != 1 {
		t.Errorf("Expected 1 winner, got %d winners", len(winners))
	}

	want := p2
	got := winners[0]
	if want != got.Player {
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

	g.Board.TableCards = []poker.Cards{c("Ks"), c("5s"), c("5c"), c("9c"), c("7h")}
	tableCards := poker.JoinCards(g.Board.TableCards...)

	winners := poker.GetWinners(tableCards, g.Players)
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

	g.Board.TableCards = []poker.Cards{c("Td"), c("Qh"), c("9s"), c("6s"), c("7h")}
	tableCards := poker.JoinCards(g.Board.TableCards...)

	winners := poker.GetWinners(tableCards, g.Players)
	if len(winners) != 1 {
		t.Errorf("Expected 1 winner, got %d winners", len(winners))
	}

	want := p1
	got := winners[0]
	if want != got.Player {
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

	g.Board.TableCards = []poker.Cards{c("Jc"), c("9h"), c("Ks"), c("8s"), c("7s")}
	tableCards := poker.JoinCards(g.Board.TableCards...)

	winners := poker.GetWinners(tableCards, g.Players)
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

	g.Board.TableCards = []poker.Cards{c("Jc"), c("9h"), c("Ks"), c("8s"), c("7s")}
	tableCards := poker.JoinCards(g.Board.TableCards...)

	winners := poker.GetWinners(tableCards, g.Players)
	if len(winners) != 1 {
		t.Errorf("Expected 1 winner, got %d", len(winners))
	}

	want := p2
	got := winners[0]
	if want != got.Player {
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

	g.Board.TableCards = []poker.Cards{c("6h"), c("Th"), c("Kh"), c("Ah"), c("5h")}
	tableCards := poker.JoinCards(g.Board.TableCards...)

	winners := poker.GetWinners(tableCards, g.Players)
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

	g.Board.TableCards = []poker.Cards{c("6h"), c("Th"), c("Kh"), c("Ad"), c("5d")}
	tableCards := poker.JoinCards(g.Board.TableCards...)

	winners := poker.GetWinners(tableCards, g.Players)
	if len(winners) != 1 {
		t.Errorf("Expected 1 winner, got %d", len(winners))
	}

	want := p1
	got := winners[0]
	if want != got.Player {
		t.Errorf("\nWant %v\nGot  %v", want, got)
	}
}

func TestFullHouseTie(t *testing.T) {
	g := poker.NewGame()
	p1 := poker.NewPlayer("P1")
	p2 := poker.NewPlayer("P2")
	g.Players = []*poker.Player{p1, p2}

	c := poker.NewCard
	p1.Hand = c("9h") | c("8h")
	p2.Hand = c("8c") | c("9c")

	g.Board.TableCards = []poker.Cards{c("8d"), c("Ts"), c("9s"), c("Ah"), c("8s")}
	tableCards := poker.JoinCards(g.Board.TableCards...)

	winners := poker.GetWinners(tableCards, g.Players)
	if len(winners) != 2 {
		t.Errorf("Expected tie, got %d winners", len(winners))
	}
}

func TestFullHouseWinsTwoPair(t *testing.T) {
	g := poker.NewGame()
	p1 := poker.NewPlayer("P1")
	p2 := poker.NewPlayer("P2")
	g.Players = []*poker.Player{p1, p2}

	c := poker.NewCard
	p1.Hand = c("Ah") | c("As")
	p2.Hand = c("Td") | c("9s")

	g.Board.TableCards = []poker.Cards{c("6s"), c("Th"), c("9h"), c("Ts"), c("5d")}
	tableCards := poker.JoinCards(g.Board.TableCards...)

	winners := poker.GetWinners(tableCards, g.Players)
	if len(winners) != 1 {
		t.Errorf("Expected 1 winner, got %d", len(winners))
	}

	want := p2
	got := winners[0]
	if want != got.Player {
		t.Errorf("\nWant %v\nGot  %v", want, got)
	}
}

func TestFourOfAKindTie(t *testing.T) {
	g := poker.NewGame()
	p1 := poker.NewPlayer("P1")
	p2 := poker.NewPlayer("P2")
	g.Players = []*poker.Player{p1, p2}

	c := poker.NewCard
	p1.Hand = c("8d") | c("8h")
	p2.Hand = c("8c") | c("3c")

	g.Board.TableCards = []poker.Cards{c("9c"), c("9d"), c("9s"), c("9h"), c("8s")}
	tableCards := poker.JoinCards(g.Board.TableCards...)

	winners := poker.GetWinners(tableCards, g.Players)
	if len(winners) != 2 {
		t.Errorf("Expected tie, got %d winners", len(winners))
	}
}

func TestFourOfAKindWinsThreeOfAKind(t *testing.T) {
	g := poker.NewGame()
	p1 := poker.NewPlayer("P1")
	p2 := poker.NewPlayer("P2")
	g.Players = []*poker.Player{p1, p2}

	c := poker.NewCard
	p1.Hand = c("2h") | c("2s")
	p2.Hand = c("Kd") | c("Ks")

	g.Board.TableCards = []poker.Cards{c("Kc"), c("Tc"), c("9c"), c("2c"), c("2d")}
	tableCards := poker.JoinCards(g.Board.TableCards...)

	winners := poker.GetWinners(tableCards, g.Players)
	if len(winners) != 1 {
		t.Errorf("Expected 1 winner, got %d", len(winners))
	}

	want := p1
	got := winners[0]
	if want != got.Player {
		t.Errorf("\nWant %v\nGot  %v", want, got)
	}
}

func TestStraightFlushTie(t *testing.T) {
	g := poker.NewGame()
	p1 := poker.NewPlayer("P1")
	p2 := poker.NewPlayer("P2")
	g.Players = []*poker.Player{p1, p2}

	c := poker.NewCard
	p1.Hand = c("As") | c("Ks")
	p2.Hand = c("3c") | c("4c")

	g.Board.TableCards = []poker.Cards{c("2s"), c("3s"), c("4s"), c("5s"), c("6s")}
	tableCards := poker.JoinCards(g.Board.TableCards...)

	winners := poker.GetWinners(tableCards, g.Players)
	if len(winners) != 2 {
		t.Errorf("Expected tie, got %d winners", len(winners))
	}
}

func TestStraightFlushWinsPair(t *testing.T) {
	g := poker.NewGame()
	p1 := poker.NewPlayer("P1")
	p2 := poker.NewPlayer("P2")
	g.Players = []*poker.Player{p1, p2}

	c := poker.NewCard
	p1.Hand = c("As") | c("7s")
	p2.Hand = c("3c") | c("9c")

	g.Board.TableCards = []poker.Cards{c("2s"), c("3s"), c("4s"), c("5s"), c("6s")}
	tableCards := poker.JoinCards(g.Board.TableCards...)

	winners := poker.GetWinners(tableCards, g.Players)
	if len(winners) != 1 {
		t.Errorf("Expected 1 winner, got %d", len(winners))
	}

	want := p1
	got := winners[0]
	if want != got.Player {
		t.Errorf("\nWant %v\nGot  %v", want, got)
	}
}

func TestRoyalFlushTie(t *testing.T) {
	g := poker.NewGame()
	p1 := poker.NewPlayer("P1")
	p2 := poker.NewPlayer("P2")
	g.Players = []*poker.Player{p1, p2}

	c := poker.NewCard
	p1.Hand = c("Ac") | c("Kc")
	p2.Hand = c("3c") | c("4c")

	g.Board.TableCards = []poker.Cards{c("Ts"), c("Ks"), c("Js"), c("Qs"), c("As")}
	tableCards := poker.JoinCards(g.Board.TableCards...)

	winners := poker.GetWinners(tableCards, g.Players)
	if len(winners) != 2 {
		t.Errorf("Expected tie, got %d winners", len(winners))
	}
}

func TestRoyalFlushWinsPair(t *testing.T) {
	g := poker.NewGame()
	p1 := poker.NewPlayer("P1")
	p2 := poker.NewPlayer("P2")
	g.Players = []*poker.Player{p1, p2}

	c := poker.NewCard
	p1.Hand = c("As") | c("7s")
	p2.Hand = c("3c") | c("9c")

	g.Board.TableCards = []poker.Cards{c("Ts"), c("Ks"), c("Js"), c("Qs"), c("9s")}
	tableCards := poker.JoinCards(g.Board.TableCards...)

	winners := poker.GetWinners(tableCards, g.Players)
	if len(winners) != 1 {
		t.Errorf("Expected 1 winner, got %d", len(winners))
	}

	want := p1
	got := winners[0]
	if want != got.Player {
		t.Errorf("\nWant %v\nGot  %v", want, got)
	}
}
