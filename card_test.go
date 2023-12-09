package poker

import (
	"testing"
)

func TestNewCard(t *testing.T) {
	want := ACES & SPADES
	got := NewCard("As")
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestQuitCards(t *testing.T) {
	c := NewCard
	cards := c("As") | c("Kc") | c("Ks") | c("3d")
	cardsToQuit := c("3d") | c("5d")

	want := c("As") | c("Kc") | c("Ks")
	got := cards.QuitCards(cardsToQuit)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestOnes(t *testing.T) {
	c := NewCard
	cards := c("Ac") | c("Kh") | c("2c") | c("3s") | c("4d") | c("5d")

	want := 6
	got := cards.Count()
	if want != got {
		t.Errorf("\nWant %d\nGot  %d", want, got)
	}
}

func TestMergeSuits(t *testing.T) {
	c := NewCard
	cards := c("As") | c("Kc") | c("Ks") | c("Jd")

	want := (ACES | KINGS | JACKS) & FIRST_SUIT
	got := cards.mergeSuits()
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestMergeSuits2(t *testing.T) {
	c := NewCard
	cards := c("Ac") | c("Kh") | c("2c") | c("3s") | c("4d") | c("5d")

	want := (ACES | KINGS | FIVES | FOURS | THREES | TWOS) & FIRST_SUIT
	got := cards.mergeSuits()
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestReduceRepeatedNumber(t *testing.T) {
	c := NewCard
	cards := c("Kc") | c("Kh") | c("Kd") | c("Ks")

	want := 2
	cards = cards.reduceRepeatedNumber(2)
	got := cards.Count()
	if want != got {
		t.Errorf("\nWant %d\nGot  %d", want, got)
	}
}

func TestQuitStraightRepeatedNumbers(t *testing.T) {
	c := NewCard
	cards := c("Kc") | c("Kh") | c("Ad") | c("Qs") | c("Td") | c("Js")

	want := 5
	cards = cards.quitStraightRepeatedNumbers()
	got := cards.Count()
	if want != got {
		t.Errorf("\nWant %d\nGot  %d", want, got)
	}
}

func TestGetJustOneFlush(t *testing.T) {
	c := NewCard
	cards := c("Ah") | c("Kh") | c("Qh") | c("Jh") | c("Th") | c("Ad")

	want := c("Ah") | c("Kh") | c("Qh") | c("Jh") | c("Th")
	got := cards.getJustOneFlush()
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestGetFlushHighestNumbers(t *testing.T) {
	c := NewCard
	cards := c("Kh") | c("Qh") | c("Jh") | c("Th") | c("9h") | c("8h") | c("7h") | c("6h") | c("5h") | c("4h")

	want := c("Kh") | c("Qh") | c("Jh") | c("Th") | c("9h")
	got := cards.getFlushHighestNumbers()
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestHighCardAce(t *testing.T) {
	c := NewCard
	cards := c("As") | c("Kc")

	want := c("As")
	got, _ := HighCard(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestHighCardSeven(t *testing.T) {
	c := NewCard
	cards := c("2s") | c("7c")

	want := c("7c")
	got, _ := HighCard(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestNoPair(t *testing.T) {
	c := NewCard
	cards := c("As") | c("Kd") | c("Qc") | c("Jc") | c("Tc") | c("9c") | c("8c")

	_, found := Pair(cards)
	if found {
		t.Errorf("\nLooking for a pair in %s\nWant found=false\nGot  found=true", cards)
	}
}

func TestPairAce(t *testing.T) {
	c := NewCard
	cards := c("As") | c("Kc") | c("Ac")

	want := c("As") | c("Ac")
	got, _ := Pair(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestPairSeven(t *testing.T) {
	c := NewCard
	cards := c("7s") | c("Kc") | c("7c")

	want := c("7s") | c("7c")
	got, _ := Pair(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestNoTwoPair(t *testing.T) {
	c := NewCard
	cards := c("Ac") | c("Ad") | c("Qc") | c("Jc") | c("Tc")

	_, found := TwoPair(cards)
	if found {
		t.Errorf("\nLooking for a two pair in %s\nWant found=false\nGot  found=true", cards)
	}
}

func TestTwoPairAceSeven(t *testing.T) {
	c := NewCard
	cards := c("7s") | c("Kc") | c("7c") | c("As") | c("Ad")

	want := c("7s") | c("7c") | c("As") | c("Ad")
	got, _ := TwoPair(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestNoThreeOfAKind(t *testing.T) {
	c := NewCard
	cards := c("Ac") | c("Ad") | c("Qc") | c("Jc") | c("Tc")

	_, found := ThreeOfAKind(cards)
	if found {
		t.Errorf("\nLooking for a three of a kind in %s\nWant found=false\nGot  found=true", cards)
	}
}

func TestThreeOfAKindAce(t *testing.T) {
	c := NewCard
	cards := c("Ah") | c("Kc") | c("7c") | c("As") | c("Ad")

	want := c("Ah") | c("As") | c("Ad")
	got, _ := ThreeOfAKind(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestNoStraight(t *testing.T) {
	c := NewCard
	cards := c("Ac") | c("Ad") | c("Qc") | c("Jc") | c("Tc")

	_, found := Straight(cards)
	if found {
		t.Errorf("\nLooking for a straight in %s\nWant found=false\nGot  found=true", cards)
	}
}

func TestStraightToFive(t *testing.T) {
	c := NewCard
	cards := c("Ac") | c("Kh") | c("2c") | c("3s") | c("4d") | c("5d")

	want := c("5d") | c("4d") | c("3s") | c("2c") | c("Ac")
	got, _ := Straight(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestStraightToTen(t *testing.T) {
	c := NewCard
	cards := c("Qh") | c("Tc") | c("9c") | c("8s") | c("7d") | c("6d")

	want := c("Tc") | c("9c") | c("8s") | c("7d") | c("6d")
	got, _ := Straight(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestStraightWithRepeatedCard(t *testing.T) {
	c := NewCard
	cards := c("8h") | c("9c") | c("8s") | c("7d") | c("6d") | c("5d")

	want := 5
	winningCards, _ := Straight(cards)
	got := winningCards.Count()
	if want != got {
		t.Errorf("\nWant %d\nGot  %d", want, got)
	}
}

func TestNoFlush(t *testing.T) {
	c := NewCard
	cards := c("Ac") | c("Ad") | c("Qc") | c("Jc") | c("Tc")

	_, found := Flush(cards)
	if found {
		t.Errorf("\nLooking for a flush in %s\nWant found=false\nGot  found=true", cards)
	}
}

func TestFlushToTen(t *testing.T) {
	c := NewCard
	cards := c("2h") | c("Th") | c("9c") | c("8h") | c("7h") | c("6h")

	want := c("2h") | c("Th") | c("8h") | c("7h") | c("6h")
	got, _ := Flush(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestFlushWithSevenSameColorCards(t *testing.T) {
	c := NewCard
	cards := c("2h") | c("Th") | c("9h") | c("8h") | c("7h") | c("6h") | c("Ah")

	want := c("Ah") | c("Th") | c("8h") | c("7h") | c("9h")
	got, _ := Flush(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestNoFullHouse(t *testing.T) {
	c := NewCard
	cards := c("Ac") | c("Ad") | c("Ah") | c("Jc") | c("Tc")

	_, found := FullHouse(cards)
	if found {
		t.Errorf("\nLooking for a full house in %s\nWant found=false\nGot  found=true", cards)
	}
}

func TestFullHouseToNine(t *testing.T) {
	c := NewCard
	cards := c("2h") | c("9h") | c("9c") | c("9d") | c("2d") | c("6h")

	want := c("2h") | c("9h") | c("9c") | c("9d") | c("2d")
	got, _ := FullHouse(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestNoFourOfAKind(t *testing.T) {
	c := NewCard
	cards := c("Ac") | c("Ad") | c("Qc") | c("Jc") | c("Tc")

	_, found := FourOfAKind(cards)
	if found {
		t.Errorf("\nLooking for a four of a kind in %s\nWant found=false\nGot  found=true", cards)
	}
}

func TestFourOfAKindAce(t *testing.T) {
	c := NewCard
	cards := c("Ah") | c("Ac") | c("Kc") | c("7c") | c("As") | c("Ad")

	want := c("Ah") | c("As") | c("Ad") | c("Ac")
	got, _ := FourOfAKind(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestNoStraightFlush(t *testing.T) {
	c := NewCard
	cards := c("Ah") | c("Ad") | c("Qh") | c("3c") | c("Tc")

	_, found := StraightFlush(cards)
	if found {
		t.Errorf("\nLooking for a straight flush in %s\nWant found=false\nGot  found=true", cards)
	}
}

func TestStraightFlushTen(t *testing.T) {
	c := NewCard
	cards := c("Th") | c("9h") | c("8h") | c("7h") | c("6h") | c("Ad")

	want := c("Th") | c("9h") | c("8h") | c("7h") | c("6h")
	got, _ := StraightFlush(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestStraightFlushFive(t *testing.T) {
	c := NewCard
	cards := c("5h") | c("4h") | c("3h") | c("2h") | c("Ah") | c("Ad")

	want := c("5h") | c("4h") | c("3h") | c("2h") | c("Ah")
	got, _ := StraightFlush(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}

func TestNoRoyalFlush(t *testing.T) {
	c := NewCard
	cards := c("Ah") | c("Ad") | c("Qh") | c("3c") | c("Tc")

	_, found := RoyalFlush(cards)
	if found {
		t.Errorf("\nLooking for a straight flush in %s\nWant found=false\nGot  found=true", cards)
	}
}

func TestRoyalFlush(t *testing.T) {
	c := NewCard
	cards := c("Th") | c("Jh") | c("Qh") | c("Kh") | c("Ah") | c("Ad")

	want := c("Th") | c("Jh") | c("Qh") | c("Kh") | c("Ah")
	got, _ := RoyalFlush(cards)
	if want != got {
		t.Errorf("\nWant %s\nGot  %s", want, got)
	}
}
