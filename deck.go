package poker

import "math/rand"

type Deck struct {
	Cards   []Cards
	pointer int
}

func NewDeck() *Deck {
	deck := &Deck{
		Cards: make([]Cards, 0, TOTAL_CARDS),
	}

	for card := Cards(0b1); card < ACES; card <<= 1 {
		deck.Cards = append(deck.Cards, card)
	}

	return deck
}

func (d *Deck) Shuffle() {
	d.pointer = 0
	for i := 0; i < TOTAL_CARDS; i++ {
		j := rand.Intn(TOTAL_CARDS)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

func (d *Deck) GetNextCard() Cards {
	if d.pointer > TOTAL_CARDS {
		return NO_CARD
	}

	card := d.Cards[d.pointer]
	d.pointer++

	return card
}
