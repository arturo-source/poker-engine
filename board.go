package poker

type BoardState int

const (
	PREFLOP BoardState = iota
	FLOP
	TURN
	RIVER
	SHOWDOWN
)

type Board struct {
	Deck        *Deck
	TableCards  []Cards
	BurnedCards []Cards
	State       BoardState
}

func NewBoard(d *Deck) *Board {
	return &Board{
		Deck:        d,
		TableCards:  make([]Cards, 0, MAX_CARDS_IN_BOARD),
		BurnedCards: make([]Cards, 0, MAX_BURNED_CARDS),
		State:       PREFLOP,
	}
}

func (b *Board) showCard() error {
	if len(b.TableCards) >= MAX_CARDS_IN_BOARD {
		return errNoCardsToFlip
	}

	card := b.Deck.GetNextCard()
	if card == NO_CARD {
		return errNoCardsInDeck
	}

	b.TableCards = append(b.TableCards, card)
	return nil
}

func (b *Board) burnCard() error {
	if len(b.BurnedCards) >= MAX_BURNED_CARDS {
		return errNoCardsToFlip
	}

	card := b.Deck.GetNextCard()
	if card == NO_CARD {
		return errNoCardsInDeck
	}

	b.BurnedCards = append(b.BurnedCards, card)
	return nil
}

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
	default:
		return errNoCardsToFlip
	}

	b.State++
	return nil
}
