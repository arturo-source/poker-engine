package poker

type BoardState int

const (
	PREFLOP BoardState = iota
	FLOP
	TURN
	RIVER
	SHOWDOWN
)

// Board represents a table where you can access to the current flipped cards, burned cards, and the board state (preflop, flop, etc.).
type Board struct {
	deck        *Deck
	TableCards  []Cards
	BurnedCards []Cards
	State       BoardState
}

// NewBoard creates a board with an specific deck, and sets the boards as the initial state.
func NewBoard(d *Deck) *Board {
	b := Board{deck: d}
	b.setInitialState()

	return &b
}

// setInitialState shuffles the cards in deck,
// creates an empty slice for the table cards,
// another one for the burned ones,
// and sets the board state to PREFLOP.
func (b *Board) setInitialState() {
	b.deck.Shuffle()
	b.TableCards = make([]Cards, 0, MAX_CARDS_IN_BOARD)
	b.BurnedCards = make([]Cards, 0, MAX_BURNED_CARDS)
	b.State = PREFLOP
}

// showCard adds the next card to the TableCards slice.
// But returns error if you exceed the max cards in board,
// or there are no more cards in the deck.
func (b *Board) showCard() error {
	if len(b.TableCards) >= MAX_CARDS_IN_BOARD {
		return errNoCardsToFlip
	}

	card := b.deck.GetNextCard()
	if card == NO_CARD {
		return errNoCardsInDeck
	}

	b.TableCards = append(b.TableCards, card)
	return nil
}

// burnCard adds the next card to the BurnedCards slice.
// But returns error if you exceed the max cards burned,
// or there are no more cards in the deck.
func (b *Board) burnCard() error {
	if len(b.BurnedCards) >= MAX_BURNED_CARDS {
		return errNoCardsToFlip
	}

	card := b.deck.GetNextCard()
	if card == NO_CARD {
		return errNoCardsInDeck
	}

	b.BurnedCards = append(b.BurnedCards, card)
	return nil
}

// Restart sets the board empty (cards burned, in table), sets the state to PREFLOP, and shuffles the deck, as the initial state.
func (b *Board) Restart() {
	b.setInitialState()
}

// NextBoardState add corresponding cards to TableCards and BurnedCards, depending on the current State.
// Returns an error if there are no more cards in deck, or if you try to get next state in the SHOWDOWN.
func (b *Board) NextBoardState() error {
	switch b.State {
	case PREFLOP:
		if err := b.burnCard(); err != nil {
			return err
		}

		for i := 0; i < 3; i++ {
			if err := b.showCard(); err != nil {
				return err
			}
		}
	case FLOP:
		if err := b.burnCard(); err != nil {
			return err
		}
		if err := b.showCard(); err != nil {
			return err
		}
	case TURN:
		if err := b.burnCard(); err != nil {
			return err
		}
		if err := b.showCard(); err != nil {
			return err
		}
	case RIVER:
		// pass to showdown
	default:
		return errNoCardsToFlip
	}

	b.State++
	return nil
}
