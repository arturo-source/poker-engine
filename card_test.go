package poker_test

import (
	"testing"

	"github.com/arturo-source/poker-engine"
)

func TestNewCard(t *testing.T) {
	want := poker.ACES & poker.SPADES
	got := poker.NewCard("As")
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestQuitCards(t *testing.T) {
	c := poker.NewCard
	cards := c("As") | c("Kc") | c("Ks") | c("3d")
	cardsToQuit := c("3d") | c("5d")

	want := c("As") | c("Kc") | c("Ks")
	got := cards.QuitCards(cardsToQuit)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestMergeSuits(t *testing.T) {
	c := poker.NewCard
	cards := c("As") | c("Kc") | c("Ks") | c("Jd")

	want := poker.Card(0b1101000000000)
	got := cards.AllSuitsToOneSuit()
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestMergeSuits2(t *testing.T) {
	c := poker.NewCard
	cards := c("Ac") | c("Kh") | c("2c") | c("3s") | c("4d") | c("5d")

	want := poker.Card(0b1100000001111)
	got := cards.AllSuitsToOneSuit()
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestHighCardAce(t *testing.T) {
	c := poker.NewCard
	cards := c("As") | c("Kc")

	want := c("As")
	got, _ := poker.HighCard(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestHighCardSeven(t *testing.T) {
	c := poker.NewCard
	cards := c("2s") | c("7c")

	want := c("7c")
	got, _ := poker.HighCard(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestNoPair(t *testing.T) {
	c := poker.NewCard
	cards := c("As") | c("Kd") | c("Qc") | c("Jc") | c("Tc") | c("9c") | c("8c")

	_, found := poker.Pair(cards)
	if found {
		t.Errorf("\nLooking for a pair in %s\nWant found=false\nGot  found=true", cards)
	}
}

func TestPairAce(t *testing.T) {
	c := poker.NewCard
	cards := c("As") | c("Kc") | c("Ac")

	want := c("As") | c("Ac")
	got, _ := poker.Pair(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestPairSeven(t *testing.T) {
	c := poker.NewCard
	cards := c("7s") | c("Kc") | c("7c")

	want := c("7s") | c("7c")
	got, _ := poker.Pair(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestNoTwoPair(t *testing.T) {
	c := poker.NewCard
	cards := c("Ac") | c("Ad") | c("Qc") | c("Jc") | c("Tc")

	_, found := poker.TwoPair(cards)
	if found {
		t.Errorf("\nLooking for a two pair in %s\nWant found=false\nGot  found=true", cards)
	}
}

func TestTwoPairAceSeven(t *testing.T) {
	c := poker.NewCard
	cards := c("7s") | c("Kc") | c("7c") | c("As") | c("Ad")

	want := c("7s") | c("7c") | c("As") | c("Ad")
	got, _ := poker.TwoPair(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestNoThreeOfAKind(t *testing.T) {
	c := poker.NewCard
	cards := c("Ac") | c("Ad") | c("Qc") | c("Jc") | c("Tc")

	_, found := poker.ThreeOfAKind(cards)
	if found {
		t.Errorf("\nLooking for a three of a kind in %s\nWant found=false\nGot  found=true", cards)
	}
}

func TestThreeOfAKindAce(t *testing.T) {
	c := poker.NewCard
	cards := c("Ah") | c("Kc") | c("7c") | c("As") | c("Ad")

	want := c("Ah") | c("As") | c("Ad")
	got, _ := poker.ThreeOfAKind(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestNoStraight(t *testing.T) {
	c := poker.NewCard
	cards := c("Ac") | c("Ad") | c("Qc") | c("Jc") | c("Tc")

	_, found := poker.Straight(cards)
	if found {
		t.Errorf("\nLooking for a straight in %s\nWant found=false\nGot  found=true", cards)
	}
}

func TestStraightToFive(t *testing.T) {
	c := poker.NewCard
	cards := c("Ac") | c("Kh") | c("2c") | c("3s") | c("4d") | c("5d")

	want := c("5d") | c("4d") | c("3s") | c("2c") | c("Ac")
	got, _ := poker.Straight(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestStraightToTen(t *testing.T) {
	c := poker.NewCard
	cards := c("Qh") | c("Tc") | c("9c") | c("8s") | c("7d") | c("6d")

	want := c("Tc") | c("9c") | c("8s") | c("7d") | c("6d")
	got, _ := poker.Straight(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestNoFlush(t *testing.T) {
	c := poker.NewCard
	cards := c("Ac") | c("Ad") | c("Qc") | c("Jc") | c("Tc")

	_, found := poker.Flush(cards)
	if found {
		t.Errorf("\nLooking for a flush in %s\nWant found=false\nGot  found=true", cards)
	}
}

func TestFlushToTen(t *testing.T) {
	c := poker.NewCard
	cards := c("2h") | c("Th") | c("9c") | c("8h") | c("7h") | c("6h")

	want := c("2h") | c("Th") | c("8h") | c("7h") | c("6h")
	got, _ := poker.Flush(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestNoFullHouse(t *testing.T) {
	c := poker.NewCard
	cards := c("Ac") | c("Ad") | c("Ah") | c("Jc") | c("Tc")

	_, found := poker.FullHouse(cards)
	if found {
		t.Errorf("\nLooking for a full house in %s\nWant found=false\nGot  found=true", cards)
	}
}

func TestFullHouseToNine(t *testing.T) {
	c := poker.NewCard
	cards := c("2h") | c("9h") | c("9c") | c("9d") | c("2d") | c("6h")

	want := c("2h") | c("9h") | c("9c") | c("9d") | c("2d")
	got, _ := poker.FullHouse(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestNoFourOfAKind(t *testing.T) {
	c := poker.NewCard
	cards := c("Ac") | c("Ad") | c("Qc") | c("Jc") | c("Tc")

	_, found := poker.FourOfAKind(cards)
	if found {
		t.Errorf("\nLooking for a four of a kind in %s\nWant found=false\nGot  found=true", cards)
	}
}

func TestFourOfAKindAce(t *testing.T) {
	c := poker.NewCard
	cards := c("Ah") | c("Ac") | c("Kc") | c("7c") | c("As") | c("Ad")

	want := c("Ah") | c("As") | c("Ad") | c("Ac")
	got, _ := poker.FourOfAKind(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestNoStraightFlush(t *testing.T) {
	c := poker.NewCard
	cards := c("Ah") | c("Ad") | c("Qh") | c("3c") | c("Tc")

	_, found := poker.StraightFlush(cards)
	if found {
		t.Errorf("\nLooking for a straight flush in %s\nWant found=false\nGot  found=true", cards)
	}
}

func TestStraightFlushTen(t *testing.T) {
	c := poker.NewCard
	cards := c("Th") | c("9h") | c("8h") | c("7h") | c("6h") | c("Ad")

	want := c("Th") | c("9h") | c("8h") | c("7h") | c("6h")
	got, _ := poker.StraightFlush(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestNoRoyalFlush(t *testing.T) {
	c := poker.NewCard
	cards := c("Ah") | c("Ad") | c("Qh") | c("3c") | c("Tc")

	_, found := poker.RoyalFlush(cards)
	if found {
		t.Errorf("\nLooking for a straight flush in %s\nWant found=false\nGot  found=true", cards)
	}
}

func TestRoyalFlush(t *testing.T) {
	c := poker.NewCard
	cards := c("Th") | c("Jh") | c("Qh") | c("Kh") | c("Ah") | c("Ad")

	want := c("Th") | c("Jh") | c("Qh") | c("Kh") | c("Ah")
	got, _ := poker.RoyalFlush(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}
