package poker

import "math/rand"

// Deck represents a deck with the 52 cards,
// you should always call NewDeck to build a deck.
type Deck struct {
	cards   []Cards
	pointer int
}

// NewDeck fills the deck with the 52 cards, and returns the reference to this deck.
func NewDeck() *Deck {
	deck := &Deck{
		cards: make([]Cards, 0, TOTAL_CARDS),
	}

	for card := Cards(0b1); card < ACES; card <<= 1 {
		deck.cards = append(deck.cards, card)
	}

	return deck
}

// Shuffle resets the pointer to 0, to start using the deck again, and shuffles the cards to get in a random order.
func (d *Deck) Shuffle() {
	d.pointer = 0
	for i := 0; i < TOTAL_CARDS; i++ {
		j := rand.Intn(TOTAL_CARDS)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

// GetNextCard returns the next card in the deck, and moves the pointer to the next card.
func (d *Deck) GetNextCard() Cards {
	if d.pointer > TOTAL_CARDS {
		return NO_CARD
	}

	card := d.cards[d.pointer]
	d.pointer++

	return card
}
