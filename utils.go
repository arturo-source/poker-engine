package poker

import "errors"

var (
	errNoCardsInDeck     = errors.New("no more cards in deck")
	errNoCardsToFlip     = errors.New("no more cards to flip")
	errMaxCardsInHand    = errors.New("max cards added to hand")
	errCardAlreadyInHand = errors.New("card already added to hand")
)

const (
	MAX_CARDS_IN_BOARD = 5
	MAX_BURNED_CARDS   = 3
	MAX_CARDS_PER_HAND = 2
)

const (
	NO_CARD   Cards = 0
	ONE_SUIT  Cards = 0b1111111111111
	ALL_CARDS Cards = 0b1111111111111111111111111111111111111111111111111111
)

const (
	CLUBS = ONE_SUIT << (iota * 13)
	DIAMONDS
	HEARTS
	SPADES
)

const (
	TWOS Cards = 0b0000000000001000000000000100000000000010000000000001 << iota
	THREES
	FOURS
	FIVES
	SIXS
	SEVENS
	EIGHTS
	NINES
	TENS
	JACKS
	QUEENS
	KINGS
	ACES
)

var (
	TOTAL_CARDS = len(SUIT_VALUES) * len(NUMBER_VALUES)

	SUIT_VALUES = map[Cards]string{
		// DIAMONDS: "♦",
		// CLUBS:    "♣",
		// HEARTS:   "♥",
		// SPADES:   "♠",
		DIAMONDS: "d",
		CLUBS:    "c",
		HEARTS:   "h",
		SPADES:   "s",
	}

	NUMBER_VALUES = map[Cards]string{
		ACES:   "A",
		KINGS:  "K",
		QUEENS: "Q",
		JACKS:  "J",
		TENS:   "T",
		NINES:  "9",
		EIGHTS: "8",
		SEVENS: "7",
		SIXS:   "6",
		FIVES:  "5",
		FOURS:  "4",
		THREES: "3",
		TWOS:   "2",
	}
)

func clamp(v, min, max int) int {
	if v > max {
		return max
	}
	if v < min {
		return min
	}

	return v
}
